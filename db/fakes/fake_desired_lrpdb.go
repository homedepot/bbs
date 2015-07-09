// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type FakeDesiredLRPDB struct {
	DesiredLRPsStub        func(lager.Logger) (*models.DesiredLRPs, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		arg1 lager.Logger
	}
	desiredLRPsReturns struct {
		result1 *models.DesiredLRPs
		result2 error
	}
}

func (fake *FakeDesiredLRPDB) DesiredLRPs(arg1 lager.Logger) (*models.DesiredLRPs, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(arg1)
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeDesiredLRPDB) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPsArgsForCall(i int) lager.Logger {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].arg1
}

func (fake *FakeDesiredLRPDB) DesiredLRPsReturns(result1 *models.DesiredLRPs, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 *models.DesiredLRPs
		result2 error
	}{result1, result2}
}

var _ db.DesiredLRPDB = new(FakeDesiredLRPDB)
