// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/director"
	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
)

type FakeDirectorFactory struct {
	NewStub        func(director.FactoryConfig, director.TaskReporter, director.FileReporter) (director.Director, error)
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		arg1 director.FactoryConfig
		arg2 director.TaskReporter
		arg3 director.FileReporter
	}
	newReturns struct {
		result1 director.Director
		result2 error
	}
	newReturnsOnCall map[int]struct {
		result1 director.Director
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirectorFactory) New(arg1 director.FactoryConfig, arg2 director.TaskReporter, arg3 director.FileReporter) (director.Director, error) {
	fake.newMutex.Lock()
	ret, specificReturn := fake.newReturnsOnCall[len(fake.newArgsForCall)]
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		arg1 director.FactoryConfig
		arg2 director.TaskReporter
		arg3 director.FileReporter
	}{arg1, arg2, arg3})
	fake.recordInvocation("New", []interface{}{arg1, arg2, arg3})
	fake.newMutex.Unlock()
	if fake.NewStub != nil {
		return fake.NewStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDirectorFactory) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeDirectorFactory) NewCalls(stub func(director.FactoryConfig, director.TaskReporter, director.FileReporter) (director.Director, error)) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = stub
}

func (fake *FakeDirectorFactory) NewArgsForCall(i int) (director.FactoryConfig, director.TaskReporter, director.FileReporter) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	argsForCall := fake.newArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDirectorFactory) NewReturns(result1 director.Director, result2 error) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 director.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeDirectorFactory) NewReturnsOnCall(i int, result1 director.Director, result2 error) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = nil
	if fake.newReturnsOnCall == nil {
		fake.newReturnsOnCall = make(map[int]struct {
			result1 director.Director
			result2 error
		})
	}
	fake.newReturnsOnCall[i] = struct {
		result1 director.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeDirectorFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDirectorFactory) recordInvocation(key string, args []interface{}) {
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

var _ boshdirector.DirectorFactory = new(FakeDirectorFactory)
