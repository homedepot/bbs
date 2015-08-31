// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type FakeLRPDB struct {
	ActualLRPGroupsStub        func(logger lager.Logger, filter models.ActualLRPFilter) ([]*models.ActualLRPGroup, *models.Error)
	actualLRPGroupsMutex       sync.RWMutex
	actualLRPGroupsArgsForCall []struct {
		logger lager.Logger
		filter models.ActualLRPFilter
	}
	actualLRPGroupsReturns struct {
		result1 []*models.ActualLRPGroup
		result2 *models.Error
	}
	ActualLRPGroupsByProcessGuidStub        func(logger lager.Logger, processGuid string) ([]*models.ActualLRPGroup, *models.Error)
	actualLRPGroupsByProcessGuidMutex       sync.RWMutex
	actualLRPGroupsByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	actualLRPGroupsByProcessGuidReturns struct {
		result1 []*models.ActualLRPGroup
		result2 *models.Error
	}
	ActualLRPGroupByProcessGuidAndIndexStub        func(logger lager.Logger, processGuid string, index int32) (*models.ActualLRPGroup, *models.Error)
	actualLRPGroupByProcessGuidAndIndexMutex       sync.RWMutex
	actualLRPGroupByProcessGuidAndIndexArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}
	actualLRPGroupByProcessGuidAndIndexReturns struct {
		result1 *models.ActualLRPGroup
		result2 *models.Error
	}
	ClaimActualLRPStub        func(logger lager.Logger, processGuid string, index int32, instanceKey *models.ActualLRPInstanceKey) *models.Error
	claimActualLRPMutex       sync.RWMutex
	claimActualLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int32
		instanceKey *models.ActualLRPInstanceKey
	}
	claimActualLRPReturns struct {
		result1 *models.Error
	}
	StartActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) *models.Error
	startActualLRPMutex       sync.RWMutex
	startActualLRPArgsForCall []struct {
		logger      lager.Logger
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		netInfo     *models.ActualLRPNetInfo
	}
	startActualLRPReturns struct {
		result1 *models.Error
	}
	CrashActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) *models.Error
	crashActualLRPMutex       sync.RWMutex
	crashActualLRPArgsForCall []struct {
		logger       lager.Logger
		key          *models.ActualLRPKey
		instanceKey  *models.ActualLRPInstanceKey
		errorMessage string
	}
	crashActualLRPReturns struct {
		result1 *models.Error
	}
	FailActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey, errorMessage string) *models.Error
	failActualLRPMutex       sync.RWMutex
	failActualLRPArgsForCall []struct {
		logger       lager.Logger
		key          *models.ActualLRPKey
		errorMessage string
	}
	failActualLRPReturns struct {
		result1 *models.Error
	}
	RemoveActualLRPStub        func(logger lager.Logger, processGuid string, index int32) *models.Error
	removeActualLRPMutex       sync.RWMutex
	removeActualLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}
	removeActualLRPReturns struct {
		result1 *models.Error
	}
	RetireActualLRPStub        func(logger lager.Logger, key *models.ActualLRPKey) *models.Error
	retireActualLRPMutex       sync.RWMutex
	retireActualLRPArgsForCall []struct {
		logger lager.Logger
		key    *models.ActualLRPKey
	}
	retireActualLRPReturns struct {
		result1 *models.Error
	}
	DesiredLRPsStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, *models.Error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPsReturns struct {
		result1 []*models.DesiredLRP
		result2 *models.Error
	}
	DesiredLRPByProcessGuidStub        func(logger lager.Logger, processGuid string) (*models.DesiredLRP, *models.Error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 *models.DesiredLRP
		result2 *models.Error
	}
	DesireLRPStub        func(logger lager.Logger, desiredLRP *models.DesiredLRP) *models.Error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 *models.Error
	}
	UpdateDesiredLRPStub        func(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) *models.Error
	updateDesiredLRPMutex       sync.RWMutex
	updateDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}
	updateDesiredLRPReturns struct {
		result1 *models.Error
	}
	RemoveDesiredLRPStub        func(logger lager.Logger, processGuid string) *models.Error
	removeDesiredLRPMutex       sync.RWMutex
	removeDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	removeDesiredLRPReturns struct {
		result1 *models.Error
	}
	ConvergeLRPsStub        func(logger lager.Logger)
	convergeLRPsMutex       sync.RWMutex
	convergeLRPsArgsForCall []struct {
		logger lager.Logger
	}
	GatherAndPruneLRPsStub        func(logger lager.Logger) (*models.ConvergenceInput, *models.Error)
	gatherAndPruneLRPsMutex       sync.RWMutex
	gatherAndPruneLRPsArgsForCall []struct {
		logger lager.Logger
	}
	gatherAndPruneLRPsReturns struct {
		result1 *models.ConvergenceInput
		result2 *models.Error
	}
}

func (fake *FakeLRPDB) ActualLRPGroups(logger lager.Logger, filter models.ActualLRPFilter) ([]*models.ActualLRPGroup, *models.Error) {
	fake.actualLRPGroupsMutex.Lock()
	fake.actualLRPGroupsArgsForCall = append(fake.actualLRPGroupsArgsForCall, struct {
		logger lager.Logger
		filter models.ActualLRPFilter
	}{logger, filter})
	fake.actualLRPGroupsMutex.Unlock()
	if fake.ActualLRPGroupsStub != nil {
		return fake.ActualLRPGroupsStub(logger, filter)
	} else {
		return fake.actualLRPGroupsReturns.result1, fake.actualLRPGroupsReturns.result2
	}
}

func (fake *FakeLRPDB) ActualLRPGroupsCallCount() int {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return len(fake.actualLRPGroupsArgsForCall)
}

func (fake *FakeLRPDB) ActualLRPGroupsArgsForCall(i int) (lager.Logger, models.ActualLRPFilter) {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return fake.actualLRPGroupsArgsForCall[i].logger, fake.actualLRPGroupsArgsForCall[i].filter
}

func (fake *FakeLRPDB) ActualLRPGroupsReturns(result1 []*models.ActualLRPGroup, result2 *models.Error) {
	fake.ActualLRPGroupsStub = nil
	fake.actualLRPGroupsReturns = struct {
		result1 []*models.ActualLRPGroup
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuid(logger lager.Logger, processGuid string) ([]*models.ActualLRPGroup, *models.Error) {
	fake.actualLRPGroupsByProcessGuidMutex.Lock()
	fake.actualLRPGroupsByProcessGuidArgsForCall = append(fake.actualLRPGroupsByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.actualLRPGroupsByProcessGuidMutex.Unlock()
	if fake.ActualLRPGroupsByProcessGuidStub != nil {
		return fake.ActualLRPGroupsByProcessGuidStub(logger, processGuid)
	} else {
		return fake.actualLRPGroupsByProcessGuidReturns.result1, fake.actualLRPGroupsByProcessGuidReturns.result2
	}
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuidCallCount() int {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPGroupsByProcessGuidArgsForCall)
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return fake.actualLRPGroupsByProcessGuidArgsForCall[i].logger, fake.actualLRPGroupsByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeLRPDB) ActualLRPGroupsByProcessGuidReturns(result1 []*models.ActualLRPGroup, result2 *models.Error) {
	fake.ActualLRPGroupsByProcessGuidStub = nil
	fake.actualLRPGroupsByProcessGuidReturns = struct {
		result1 []*models.ActualLRPGroup
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndex(logger lager.Logger, processGuid string, index int32) (*models.ActualLRPGroup, *models.Error) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Lock()
	fake.actualLRPGroupByProcessGuidAndIndexArgsForCall = append(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}{logger, processGuid, index})
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Unlock()
	if fake.ActualLRPGroupByProcessGuidAndIndexStub != nil {
		return fake.ActualLRPGroupByProcessGuidAndIndexStub(logger, processGuid, index)
	} else {
		return fake.actualLRPGroupByProcessGuidAndIndexReturns.result1, fake.actualLRPGroupByProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndexCallCount() int {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return len(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall)
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndexArgsForCall(i int) (lager.Logger, string, int32) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].logger, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].processGuid, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].index
}

func (fake *FakeLRPDB) ActualLRPGroupByProcessGuidAndIndexReturns(result1 *models.ActualLRPGroup, result2 *models.Error) {
	fake.ActualLRPGroupByProcessGuidAndIndexStub = nil
	fake.actualLRPGroupByProcessGuidAndIndexReturns = struct {
		result1 *models.ActualLRPGroup
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeLRPDB) ClaimActualLRP(logger lager.Logger, processGuid string, index int32, instanceKey *models.ActualLRPInstanceKey) *models.Error {
	fake.claimActualLRPMutex.Lock()
	fake.claimActualLRPArgsForCall = append(fake.claimActualLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int32
		instanceKey *models.ActualLRPInstanceKey
	}{logger, processGuid, index, instanceKey})
	fake.claimActualLRPMutex.Unlock()
	if fake.ClaimActualLRPStub != nil {
		return fake.ClaimActualLRPStub(logger, processGuid, index, instanceKey)
	} else {
		return fake.claimActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) ClaimActualLRPCallCount() int {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return len(fake.claimActualLRPArgsForCall)
}

func (fake *FakeLRPDB) ClaimActualLRPArgsForCall(i int) (lager.Logger, string, int32, *models.ActualLRPInstanceKey) {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return fake.claimActualLRPArgsForCall[i].logger, fake.claimActualLRPArgsForCall[i].processGuid, fake.claimActualLRPArgsForCall[i].index, fake.claimActualLRPArgsForCall[i].instanceKey
}

func (fake *FakeLRPDB) ClaimActualLRPReturns(result1 *models.Error) {
	fake.ClaimActualLRPStub = nil
	fake.claimActualLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) StartActualLRP(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) *models.Error {
	fake.startActualLRPMutex.Lock()
	fake.startActualLRPArgsForCall = append(fake.startActualLRPArgsForCall, struct {
		logger      lager.Logger
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		netInfo     *models.ActualLRPNetInfo
	}{logger, key, instanceKey, netInfo})
	fake.startActualLRPMutex.Unlock()
	if fake.StartActualLRPStub != nil {
		return fake.StartActualLRPStub(logger, key, instanceKey, netInfo)
	} else {
		return fake.startActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) StartActualLRPCallCount() int {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return len(fake.startActualLRPArgsForCall)
}

func (fake *FakeLRPDB) StartActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return fake.startActualLRPArgsForCall[i].logger, fake.startActualLRPArgsForCall[i].key, fake.startActualLRPArgsForCall[i].instanceKey, fake.startActualLRPArgsForCall[i].netInfo
}

func (fake *FakeLRPDB) StartActualLRPReturns(result1 *models.Error) {
	fake.StartActualLRPStub = nil
	fake.startActualLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) CrashActualLRP(logger lager.Logger, key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) *models.Error {
	fake.crashActualLRPMutex.Lock()
	fake.crashActualLRPArgsForCall = append(fake.crashActualLRPArgsForCall, struct {
		logger       lager.Logger
		key          *models.ActualLRPKey
		instanceKey  *models.ActualLRPInstanceKey
		errorMessage string
	}{logger, key, instanceKey, errorMessage})
	fake.crashActualLRPMutex.Unlock()
	if fake.CrashActualLRPStub != nil {
		return fake.CrashActualLRPStub(logger, key, instanceKey, errorMessage)
	} else {
		return fake.crashActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) CrashActualLRPCallCount() int {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return len(fake.crashActualLRPArgsForCall)
}

func (fake *FakeLRPDB) CrashActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return fake.crashActualLRPArgsForCall[i].logger, fake.crashActualLRPArgsForCall[i].key, fake.crashActualLRPArgsForCall[i].instanceKey, fake.crashActualLRPArgsForCall[i].errorMessage
}

func (fake *FakeLRPDB) CrashActualLRPReturns(result1 *models.Error) {
	fake.CrashActualLRPStub = nil
	fake.crashActualLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) FailActualLRP(logger lager.Logger, key *models.ActualLRPKey, errorMessage string) *models.Error {
	fake.failActualLRPMutex.Lock()
	fake.failActualLRPArgsForCall = append(fake.failActualLRPArgsForCall, struct {
		logger       lager.Logger
		key          *models.ActualLRPKey
		errorMessage string
	}{logger, key, errorMessage})
	fake.failActualLRPMutex.Unlock()
	if fake.FailActualLRPStub != nil {
		return fake.FailActualLRPStub(logger, key, errorMessage)
	} else {
		return fake.failActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) FailActualLRPCallCount() int {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return len(fake.failActualLRPArgsForCall)
}

func (fake *FakeLRPDB) FailActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, string) {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return fake.failActualLRPArgsForCall[i].logger, fake.failActualLRPArgsForCall[i].key, fake.failActualLRPArgsForCall[i].errorMessage
}

func (fake *FakeLRPDB) FailActualLRPReturns(result1 *models.Error) {
	fake.FailActualLRPStub = nil
	fake.failActualLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) RemoveActualLRP(logger lager.Logger, processGuid string, index int32) *models.Error {
	fake.removeActualLRPMutex.Lock()
	fake.removeActualLRPArgsForCall = append(fake.removeActualLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int32
	}{logger, processGuid, index})
	fake.removeActualLRPMutex.Unlock()
	if fake.RemoveActualLRPStub != nil {
		return fake.RemoveActualLRPStub(logger, processGuid, index)
	} else {
		return fake.removeActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) RemoveActualLRPCallCount() int {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return len(fake.removeActualLRPArgsForCall)
}

func (fake *FakeLRPDB) RemoveActualLRPArgsForCall(i int) (lager.Logger, string, int32) {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return fake.removeActualLRPArgsForCall[i].logger, fake.removeActualLRPArgsForCall[i].processGuid, fake.removeActualLRPArgsForCall[i].index
}

func (fake *FakeLRPDB) RemoveActualLRPReturns(result1 *models.Error) {
	fake.RemoveActualLRPStub = nil
	fake.removeActualLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) RetireActualLRP(logger lager.Logger, key *models.ActualLRPKey) *models.Error {
	fake.retireActualLRPMutex.Lock()
	fake.retireActualLRPArgsForCall = append(fake.retireActualLRPArgsForCall, struct {
		logger lager.Logger
		key    *models.ActualLRPKey
	}{logger, key})
	fake.retireActualLRPMutex.Unlock()
	if fake.RetireActualLRPStub != nil {
		return fake.RetireActualLRPStub(logger, key)
	} else {
		return fake.retireActualLRPReturns.result1
	}
}

func (fake *FakeLRPDB) RetireActualLRPCallCount() int {
	fake.retireActualLRPMutex.RLock()
	defer fake.retireActualLRPMutex.RUnlock()
	return len(fake.retireActualLRPArgsForCall)
}

func (fake *FakeLRPDB) RetireActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey) {
	fake.retireActualLRPMutex.RLock()
	defer fake.retireActualLRPMutex.RUnlock()
	return fake.retireActualLRPArgsForCall[i].logger, fake.retireActualLRPArgsForCall[i].key
}

func (fake *FakeLRPDB) RetireActualLRPReturns(result1 *models.Error) {
	fake.RetireActualLRPStub = nil
	fake.retireActualLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) DesiredLRPs(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, *models.Error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(logger, filter)
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeLRPDB) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeLRPDB) DesiredLRPsArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].logger, fake.desiredLRPsArgsForCall[i].filter
}

func (fake *FakeLRPDB) DesiredLRPsReturns(result1 []*models.DesiredLRP, result2 *models.Error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []*models.DesiredLRP
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuid(logger lager.Logger, processGuid string) (*models.DesiredLRP, *models.Error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(logger, processGuid)
	} else {
		return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
	}
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].logger, fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeLRPDB) DesiredLRPByProcessGuidReturns(result1 *models.DesiredLRP, result2 *models.Error) {
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 *models.DesiredLRP
		result2 *models.Error
	}{result1, result2}
}

func (fake *FakeLRPDB) DesireLRP(logger lager.Logger, desiredLRP *models.DesiredLRP) *models.Error {
	fake.desireLRPMutex.Lock()
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}{logger, desiredLRP})
	fake.desireLRPMutex.Unlock()
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(logger, desiredLRP)
	} else {
		return fake.desireLRPReturns.result1
	}
}

func (fake *FakeLRPDB) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeLRPDB) DesireLRPArgsForCall(i int) (lager.Logger, *models.DesiredLRP) {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return fake.desireLRPArgsForCall[i].logger, fake.desireLRPArgsForCall[i].desiredLRP
}

func (fake *FakeLRPDB) DesireLRPReturns(result1 *models.Error) {
	fake.DesireLRPStub = nil
	fake.desireLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) UpdateDesiredLRP(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) *models.Error {
	fake.updateDesiredLRPMutex.Lock()
	fake.updateDesiredLRPArgsForCall = append(fake.updateDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}{logger, processGuid, update})
	fake.updateDesiredLRPMutex.Unlock()
	if fake.UpdateDesiredLRPStub != nil {
		return fake.UpdateDesiredLRPStub(logger, processGuid, update)
	} else {
		return fake.updateDesiredLRPReturns.result1
	}
}

func (fake *FakeLRPDB) UpdateDesiredLRPCallCount() int {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return len(fake.updateDesiredLRPArgsForCall)
}

func (fake *FakeLRPDB) UpdateDesiredLRPArgsForCall(i int) (lager.Logger, string, *models.DesiredLRPUpdate) {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return fake.updateDesiredLRPArgsForCall[i].logger, fake.updateDesiredLRPArgsForCall[i].processGuid, fake.updateDesiredLRPArgsForCall[i].update
}

func (fake *FakeLRPDB) UpdateDesiredLRPReturns(result1 *models.Error) {
	fake.UpdateDesiredLRPStub = nil
	fake.updateDesiredLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) RemoveDesiredLRP(logger lager.Logger, processGuid string) *models.Error {
	fake.removeDesiredLRPMutex.Lock()
	fake.removeDesiredLRPArgsForCall = append(fake.removeDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.removeDesiredLRPMutex.Unlock()
	if fake.RemoveDesiredLRPStub != nil {
		return fake.RemoveDesiredLRPStub(logger, processGuid)
	} else {
		return fake.removeDesiredLRPReturns.result1
	}
}

func (fake *FakeLRPDB) RemoveDesiredLRPCallCount() int {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return len(fake.removeDesiredLRPArgsForCall)
}

func (fake *FakeLRPDB) RemoveDesiredLRPArgsForCall(i int) (lager.Logger, string) {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return fake.removeDesiredLRPArgsForCall[i].logger, fake.removeDesiredLRPArgsForCall[i].processGuid
}

func (fake *FakeLRPDB) RemoveDesiredLRPReturns(result1 *models.Error) {
	fake.RemoveDesiredLRPStub = nil
	fake.removeDesiredLRPReturns = struct {
		result1 *models.Error
	}{result1}
}

func (fake *FakeLRPDB) ConvergeLRPs(logger lager.Logger) {
	fake.convergeLRPsMutex.Lock()
	fake.convergeLRPsArgsForCall = append(fake.convergeLRPsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.convergeLRPsMutex.Unlock()
	if fake.ConvergeLRPsStub != nil {
		fake.ConvergeLRPsStub(logger)
	}
}

func (fake *FakeLRPDB) ConvergeLRPsCallCount() int {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return len(fake.convergeLRPsArgsForCall)
}

func (fake *FakeLRPDB) ConvergeLRPsArgsForCall(i int) lager.Logger {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return fake.convergeLRPsArgsForCall[i].logger
}

func (fake *FakeLRPDB) GatherAndPruneLRPs(logger lager.Logger) (*models.ConvergenceInput, *models.Error) {
	fake.gatherAndPruneLRPsMutex.Lock()
	fake.gatherAndPruneLRPsArgsForCall = append(fake.gatherAndPruneLRPsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.gatherAndPruneLRPsMutex.Unlock()
	if fake.GatherAndPruneLRPsStub != nil {
		return fake.GatherAndPruneLRPsStub(logger)
	} else {
		return fake.gatherAndPruneLRPsReturns.result1, fake.gatherAndPruneLRPsReturns.result2
	}
}

func (fake *FakeLRPDB) GatherAndPruneLRPsCallCount() int {
	fake.gatherAndPruneLRPsMutex.RLock()
	defer fake.gatherAndPruneLRPsMutex.RUnlock()
	return len(fake.gatherAndPruneLRPsArgsForCall)
}

func (fake *FakeLRPDB) GatherAndPruneLRPsArgsForCall(i int) lager.Logger {
	fake.gatherAndPruneLRPsMutex.RLock()
	defer fake.gatherAndPruneLRPsMutex.RUnlock()
	return fake.gatherAndPruneLRPsArgsForCall[i].logger
}

func (fake *FakeLRPDB) GatherAndPruneLRPsReturns(result1 *models.ConvergenceInput, result2 *models.Error) {
	fake.GatherAndPruneLRPsStub = nil
	fake.gatherAndPruneLRPsReturns = struct {
		result1 *models.ConvergenceInput
		result2 *models.Error
	}{result1, result2}
}

var _ db.LRPDB = new(FakeLRPDB)
