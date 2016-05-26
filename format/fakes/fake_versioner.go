// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/format"
)

type FakeVersioner struct {
	VersionStub        func() format.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 format.Version
	}
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct{}
	validateReturns     struct {
		result1 error
	}
}

func (fake *FakeVersioner) Version() format.Version {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
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

func (fake *FakeVersioner) Validate() error {
	fake.validateMutex.Lock()
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct{}{})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub()
	} else {
		return fake.validateReturns.result1
	}
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

var _ format.Versioner = new(FakeVersioner)
