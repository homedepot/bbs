// Code generated by counterfeiter. DO NOT EDIT.
package formatfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/format"
)

type FakeVersioner struct {
	VersionStub        func() format.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 format.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 format.Version
	}
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct{}
	validateReturns     struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVersioner) Version() format.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeVersioner) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeVersioner) VersionReturns(result1 format.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 format.Version
	}{result1}
}

func (fake *FakeVersioner) VersionReturnsOnCall(i int, result1 format.Version) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 format.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 format.Version
	}{result1}
}

func (fake *FakeVersioner) Validate() error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct{}{})
	fake.recordInvocation("Validate", []interface{}{})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *FakeVersioner) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeVersioner) ValidateReturns(result1 error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersioner) ValidateReturnsOnCall(i int, result1 error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersioner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVersioner) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ format.Versioner = new(FakeVersioner)
