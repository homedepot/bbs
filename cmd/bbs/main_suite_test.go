package main_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/grpc/grpclog"

	"code.cloudfoundry.org/bbs"
	bbsconfig "code.cloudfoundry.org/bbs/cmd/bbs/config"
	"code.cloudfoundry.org/bbs/db/etcd"
	"code.cloudfoundry.org/bbs/encryption"
	"code.cloudfoundry.org/bbs/test_helpers"
	"code.cloudfoundry.org/bbs/test_helpers/sqlrunner"
	"code.cloudfoundry.org/consuladapter"
	"code.cloudfoundry.org/consuladapter/consulrunner"
	"code.cloudfoundry.org/durationjson"
	"code.cloudfoundry.org/inigo/helpers/portauthority"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/cloudfoundry/storeadapter/storerunner/etcdstorerunner"
	etcdclient "github.com/coreos/go-etcd/etcd"
	"github.com/gogo/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"

	"testing"
	"time"
)

var (
	etcdPort    int
	etcdUrl     string
	etcdRunner  *etcdstorerunner.ETCDClusterRunner
	etcdClient  *etcdclient.Client
	storeClient etcd.StoreClient

	logger        lager.Logger
	portAllocator portauthority.PortAllocator

	client              bbs.InternalClient
	bbsBinPath          string
	bbsAddress          string
	bbsHealthAddress    string
	bbsPort             uint16
	bbsURL              *url.URL
	bbsConfig           bbsconfig.BBSConfig
	bbsRunner           *ginkgomon.Runner
	bbsProcess          ifrit.Process
	consulRunner        *consulrunner.ClusterRunner
	consulClient        consuladapter.Client
	consulHelper        *test_helpers.ConsulHelper
	auctioneerServer    *ghttp.Server
	testMetricsListener net.PacketConn
	testMetricsChan     chan *events.Envelope
	locketBinPath       string

	sqlProcess ifrit.Process
	sqlRunner  sqlrunner.SQLRunner
)

func TestBBS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BBS Cmd Suite")
}

var _ = SynchronizedBeforeSuite(
	func() []byte {
		bbsPath, err := gexec.Build("code.cloudfoundry.org/bbs/cmd/bbs", "-race")
		Expect(err).NotTo(HaveOccurred())

		locketPath, err := gexec.Build("code.cloudfoundry.org/locket/cmd/locket", "-race")
		Expect(err).NotTo(HaveOccurred())

		return []byte(strings.Join([]string{bbsPath, locketPath}, ","))
	},
	func(binPaths []byte) {
		grpclog.SetLogger(log.New(ioutil.Discard, "", 0))
		startPort := 1050 // This could be lowered when we remove etcd
		portRange := 1000
		var err error
		portAllocator, err = portauthority.New(startPort+(portRange*(GinkgoParallelNode()-1)), startPort+(GinkgoParallelNode()*portRange))
		Expect(err).NotTo(HaveOccurred())

		path := string(binPaths)
		bbsBinPath = strings.Split(path, ",")[0]
		locketBinPath = strings.Split(path, ",")[1]

		SetDefaultEventuallyTimeout(15 * time.Second)

		// The etcd cluster runner uses a port and that port+3000. We can't edit the
		// cluster runner because it's in the attic. Thus, we do not use the portallocator
		// for etcd but instead block off ports 4000-7050 for etcd. (This allows for 50
		// parallel nodes).
		//
		// See the cluser runner
		// https://github.com/cloudfoundry-attic/storeadapter/blob/master/storerunner/etcdstorerunner/etcd_cluster_runner.go#L348-L350
		etcdPort = 4001 + GinkgoParallelNode()
		etcdUrl = fmt.Sprintf("http://127.0.0.1:%d", etcdPort)
		etcdRunner = etcdstorerunner.NewETCDClusterRunner(etcdPort, 1, nil)

		dbName := fmt.Sprintf("diego_%d", GinkgoParallelNode())
		sqlRunner = test_helpers.NewSQLRunner(dbName)
		sqlProcess = ginkgomon.Invoke(sqlRunner)

		consulStartingPort, err := portAllocator.ClaimPorts(consulrunner.PortOffsetLength)
		Expect(err).NotTo(HaveOccurred())

		consulRunner = consulrunner.NewClusterRunner(
			consulrunner.ClusterRunnerConfig{
				StartingPort: int(consulStartingPort),
				NumNodes:     1,
				Scheme:       "http",
			},
		)

		consulRunner.Start()
		consulRunner.WaitUntilReady()

		etcdRunner.Start()
	},
)

var _ = SynchronizedAfterSuite(func() {
	ginkgomon.Kill(sqlProcess)

	if etcdRunner != nil {
		etcdRunner.Stop()
	}
	if consulRunner != nil {
		consulRunner.Stop()
	}
}, func() {
	gexec.CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	var err error
	logger = lagertest.NewTestLogger("test")

	etcdRunner.Reset()

	consulRunner.Reset()
	consulClient = consulRunner.NewClient()

	etcdClient = etcdRunner.Client()
	etcdClient.SetConsistency(etcdclient.STRONG_CONSISTENCY)

	auctioneerServer = ghttp.NewServer()
	auctioneerServer.UnhandledRequestStatusCode = http.StatusAccepted
	auctioneerServer.AllowUnhandledRequests = true

	bbsPort, err = portAllocator.ClaimPorts(1)
	Expect(err).NotTo(HaveOccurred())
	bbsAddress = fmt.Sprintf("127.0.0.1:%d", bbsPort)

	bbsHealthPort, err := portAllocator.ClaimPorts(1)
	Expect(err).NotTo(HaveOccurred())
	bbsHealthAddress = fmt.Sprintf("127.0.0.1:%d", bbsHealthPort)

	bbsURL = &url.URL{
		Scheme: "http",
		Host:   bbsAddress,
	}

	testMetricsListener, _ = net.ListenPacket("udp", "127.0.0.1:0")
	testMetricsChan = make(chan *events.Envelope, 1024)
	go func() {
		defer GinkgoRecover()
		defer close(testMetricsChan)

		for {
			buffer := make([]byte, 1024)
			n, _, err := testMetricsListener.ReadFrom(buffer)
			if err != nil {
				return
			}

			var envelope events.Envelope
			err = proto.Unmarshal(buffer[:n], &envelope)
			Expect(err).NotTo(HaveOccurred())
			testMetricsChan <- &envelope
		}
	}()

	port, err := strconv.Atoi(strings.TrimPrefix(testMetricsListener.LocalAddr().String(), "127.0.0.1:"))
	Expect(err).NotTo(HaveOccurred())

	client = bbs.NewClient(bbsURL.String())

	bbsConfig = bbsconfig.BBSConfig{
		ListenAddress:     bbsAddress,
		AdvertiseURL:      bbsURL.String(),
		AuctioneerAddress: auctioneerServer.URL(),
		ConsulCluster:     consulRunner.ConsulCluster(),
		DropsondePort:     port,
		ETCDConfig: bbsconfig.ETCDConfig{
			ClusterUrls: []string{etcdUrl}, // etcd is still being used to test version migration in migration_version_test.go
		},
		DatabaseDriver:                sqlRunner.DriverName(),
		DatabaseConnectionString:      sqlRunner.ConnectionString(),
		DetectConsulCellRegistrations: true,
		ReportInterval:                durationjson.Duration(10 * time.Millisecond),
		HealthAddress:                 bbsHealthAddress,

		EncryptionConfig: encryption.EncryptionConfig{
			EncryptionKeys: map[string]string{"label": "key"},
			ActiveKeyLabel: "label",
		},
		ConvergeRepeatInterval: durationjson.Duration(time.Hour),
		UUID: "bbs-bosh-boshy-bosh-bosh",
	}
	storeClient = etcd.NewStoreClient(etcdClient)
	consulHelper = test_helpers.NewConsulHelper(logger, consulClient)
})

var _ = AfterEach(func() {
	ginkgomon.Kill(bbsProcess)

	// Make sure the healthcheck server is really gone before trying to start up
	// the bbs again in another test.
	Eventually(func() error {
		conn, err := net.Dial("tcp", bbsHealthAddress)
		if err == nil {
			conn.Close()
			return nil
		}

		return err
	}).Should(HaveOccurred())

	auctioneerServer.Close()
	testMetricsListener.Close()
	Eventually(testMetricsChan).Should(BeClosed())

	sqlRunner.Reset()
})
