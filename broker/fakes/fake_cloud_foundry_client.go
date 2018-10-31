// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	log "log"
	sync "sync"

	broker "github.com/pivotal-cf/on-demand-service-broker/broker"
	cf "github.com/pivotal-cf/on-demand-service-broker/cf"
	service "github.com/pivotal-cf/on-demand-service-broker/service"
)

type FakeCloudFoundryClient struct {
	CountInstancesOfPlanStub        func(string, string, *log.Logger) (int, error)
	countInstancesOfPlanMutex       sync.RWMutex
	countInstancesOfPlanArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}
	countInstancesOfPlanReturns struct {
		result1 int
		result2 error
	}
	countInstancesOfPlanReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	CountInstancesOfServiceOfferingStub        func(string, *log.Logger) (map[cf.ServicePlan]int, error)
	countInstancesOfServiceOfferingMutex       sync.RWMutex
	countInstancesOfServiceOfferingArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	countInstancesOfServiceOfferingReturns struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}
	countInstancesOfServiceOfferingReturnsOnCall map[int]struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}
	GetAPIVersionStub        func(*log.Logger) (string, error)
	getAPIVersionMutex       sync.RWMutex
	getAPIVersionArgsForCall []struct {
		arg1 *log.Logger
	}
	getAPIVersionReturns struct {
		result1 string
		result2 error
	}
	getAPIVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetInstanceStateStub        func(string, *log.Logger) (cf.InstanceState, error)
	getInstanceStateMutex       sync.RWMutex
	getInstanceStateArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	getInstanceStateReturns struct {
		result1 cf.InstanceState
		result2 error
	}
	getInstanceStateReturnsOnCall map[int]struct {
		result1 cf.InstanceState
		result2 error
	}
	GetInstancesOfServiceOfferingStub        func(string, *log.Logger) ([]service.Instance, error)
	getInstancesOfServiceOfferingMutex       sync.RWMutex
	getInstancesOfServiceOfferingArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	getInstancesOfServiceOfferingReturns struct {
		result1 []service.Instance
		result2 error
	}
	getInstancesOfServiceOfferingReturnsOnCall map[int]struct {
		result1 []service.Instance
		result2 error
	}
	GetInstancesOfServiceOfferingByOrgSpaceStub        func(string, string, string, *log.Logger) ([]service.Instance, error)
	getInstancesOfServiceOfferingByOrgSpaceMutex       sync.RWMutex
	getInstancesOfServiceOfferingByOrgSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 *log.Logger
	}
	getInstancesOfServiceOfferingByOrgSpaceReturns struct {
		result1 []service.Instance
		result2 error
	}
	getInstancesOfServiceOfferingByOrgSpaceReturnsOnCall map[int]struct {
		result1 []service.Instance
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudFoundryClient) CountInstancesOfPlan(arg1 string, arg2 string, arg3 *log.Logger) (int, error) {
	fake.countInstancesOfPlanMutex.Lock()
	ret, specificReturn := fake.countInstancesOfPlanReturnsOnCall[len(fake.countInstancesOfPlanArgsForCall)]
	fake.countInstancesOfPlanArgsForCall = append(fake.countInstancesOfPlanArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("CountInstancesOfPlan", []interface{}{arg1, arg2, arg3})
	fake.countInstancesOfPlanMutex.Unlock()
	if fake.CountInstancesOfPlanStub != nil {
		return fake.CountInstancesOfPlanStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.countInstancesOfPlanReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) CountInstancesOfPlanCallCount() int {
	fake.countInstancesOfPlanMutex.RLock()
	defer fake.countInstancesOfPlanMutex.RUnlock()
	return len(fake.countInstancesOfPlanArgsForCall)
}

func (fake *FakeCloudFoundryClient) CountInstancesOfPlanCalls(stub func(string, string, *log.Logger) (int, error)) {
	fake.countInstancesOfPlanMutex.Lock()
	defer fake.countInstancesOfPlanMutex.Unlock()
	fake.CountInstancesOfPlanStub = stub
}

func (fake *FakeCloudFoundryClient) CountInstancesOfPlanArgsForCall(i int) (string, string, *log.Logger) {
	fake.countInstancesOfPlanMutex.RLock()
	defer fake.countInstancesOfPlanMutex.RUnlock()
	argsForCall := fake.countInstancesOfPlanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCloudFoundryClient) CountInstancesOfPlanReturns(result1 int, result2 error) {
	fake.countInstancesOfPlanMutex.Lock()
	defer fake.countInstancesOfPlanMutex.Unlock()
	fake.CountInstancesOfPlanStub = nil
	fake.countInstancesOfPlanReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) CountInstancesOfPlanReturnsOnCall(i int, result1 int, result2 error) {
	fake.countInstancesOfPlanMutex.Lock()
	defer fake.countInstancesOfPlanMutex.Unlock()
	fake.CountInstancesOfPlanStub = nil
	if fake.countInstancesOfPlanReturnsOnCall == nil {
		fake.countInstancesOfPlanReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countInstancesOfPlanReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) CountInstancesOfServiceOffering(arg1 string, arg2 *log.Logger) (map[cf.ServicePlan]int, error) {
	fake.countInstancesOfServiceOfferingMutex.Lock()
	ret, specificReturn := fake.countInstancesOfServiceOfferingReturnsOnCall[len(fake.countInstancesOfServiceOfferingArgsForCall)]
	fake.countInstancesOfServiceOfferingArgsForCall = append(fake.countInstancesOfServiceOfferingArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("CountInstancesOfServiceOffering", []interface{}{arg1, arg2})
	fake.countInstancesOfServiceOfferingMutex.Unlock()
	if fake.CountInstancesOfServiceOfferingStub != nil {
		return fake.CountInstancesOfServiceOfferingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.countInstancesOfServiceOfferingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) CountInstancesOfServiceOfferingCallCount() int {
	fake.countInstancesOfServiceOfferingMutex.RLock()
	defer fake.countInstancesOfServiceOfferingMutex.RUnlock()
	return len(fake.countInstancesOfServiceOfferingArgsForCall)
}

func (fake *FakeCloudFoundryClient) CountInstancesOfServiceOfferingCalls(stub func(string, *log.Logger) (map[cf.ServicePlan]int, error)) {
	fake.countInstancesOfServiceOfferingMutex.Lock()
	defer fake.countInstancesOfServiceOfferingMutex.Unlock()
	fake.CountInstancesOfServiceOfferingStub = stub
}

func (fake *FakeCloudFoundryClient) CountInstancesOfServiceOfferingArgsForCall(i int) (string, *log.Logger) {
	fake.countInstancesOfServiceOfferingMutex.RLock()
	defer fake.countInstancesOfServiceOfferingMutex.RUnlock()
	argsForCall := fake.countInstancesOfServiceOfferingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCloudFoundryClient) CountInstancesOfServiceOfferingReturns(result1 map[cf.ServicePlan]int, result2 error) {
	fake.countInstancesOfServiceOfferingMutex.Lock()
	defer fake.countInstancesOfServiceOfferingMutex.Unlock()
	fake.CountInstancesOfServiceOfferingStub = nil
	fake.countInstancesOfServiceOfferingReturns = struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) CountInstancesOfServiceOfferingReturnsOnCall(i int, result1 map[cf.ServicePlan]int, result2 error) {
	fake.countInstancesOfServiceOfferingMutex.Lock()
	defer fake.countInstancesOfServiceOfferingMutex.Unlock()
	fake.CountInstancesOfServiceOfferingStub = nil
	if fake.countInstancesOfServiceOfferingReturnsOnCall == nil {
		fake.countInstancesOfServiceOfferingReturnsOnCall = make(map[int]struct {
			result1 map[cf.ServicePlan]int
			result2 error
		})
	}
	fake.countInstancesOfServiceOfferingReturnsOnCall[i] = struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetAPIVersion(arg1 *log.Logger) (string, error) {
	fake.getAPIVersionMutex.Lock()
	ret, specificReturn := fake.getAPIVersionReturnsOnCall[len(fake.getAPIVersionArgsForCall)]
	fake.getAPIVersionArgsForCall = append(fake.getAPIVersionArgsForCall, struct {
		arg1 *log.Logger
	}{arg1})
	fake.recordInvocation("GetAPIVersion", []interface{}{arg1})
	fake.getAPIVersionMutex.Unlock()
	if fake.GetAPIVersionStub != nil {
		return fake.GetAPIVersionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAPIVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) GetAPIVersionCallCount() int {
	fake.getAPIVersionMutex.RLock()
	defer fake.getAPIVersionMutex.RUnlock()
	return len(fake.getAPIVersionArgsForCall)
}

func (fake *FakeCloudFoundryClient) GetAPIVersionCalls(stub func(*log.Logger) (string, error)) {
	fake.getAPIVersionMutex.Lock()
	defer fake.getAPIVersionMutex.Unlock()
	fake.GetAPIVersionStub = stub
}

func (fake *FakeCloudFoundryClient) GetAPIVersionArgsForCall(i int) *log.Logger {
	fake.getAPIVersionMutex.RLock()
	defer fake.getAPIVersionMutex.RUnlock()
	argsForCall := fake.getAPIVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudFoundryClient) GetAPIVersionReturns(result1 string, result2 error) {
	fake.getAPIVersionMutex.Lock()
	defer fake.getAPIVersionMutex.Unlock()
	fake.GetAPIVersionStub = nil
	fake.getAPIVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetAPIVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.getAPIVersionMutex.Lock()
	defer fake.getAPIVersionMutex.Unlock()
	fake.GetAPIVersionStub = nil
	if fake.getAPIVersionReturnsOnCall == nil {
		fake.getAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getAPIVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetInstanceState(arg1 string, arg2 *log.Logger) (cf.InstanceState, error) {
	fake.getInstanceStateMutex.Lock()
	ret, specificReturn := fake.getInstanceStateReturnsOnCall[len(fake.getInstanceStateArgsForCall)]
	fake.getInstanceStateArgsForCall = append(fake.getInstanceStateArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetInstanceState", []interface{}{arg1, arg2})
	fake.getInstanceStateMutex.Unlock()
	if fake.GetInstanceStateStub != nil {
		return fake.GetInstanceStateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstanceStateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) GetInstanceStateCallCount() int {
	fake.getInstanceStateMutex.RLock()
	defer fake.getInstanceStateMutex.RUnlock()
	return len(fake.getInstanceStateArgsForCall)
}

func (fake *FakeCloudFoundryClient) GetInstanceStateCalls(stub func(string, *log.Logger) (cf.InstanceState, error)) {
	fake.getInstanceStateMutex.Lock()
	defer fake.getInstanceStateMutex.Unlock()
	fake.GetInstanceStateStub = stub
}

func (fake *FakeCloudFoundryClient) GetInstanceStateArgsForCall(i int) (string, *log.Logger) {
	fake.getInstanceStateMutex.RLock()
	defer fake.getInstanceStateMutex.RUnlock()
	argsForCall := fake.getInstanceStateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCloudFoundryClient) GetInstanceStateReturns(result1 cf.InstanceState, result2 error) {
	fake.getInstanceStateMutex.Lock()
	defer fake.getInstanceStateMutex.Unlock()
	fake.GetInstanceStateStub = nil
	fake.getInstanceStateReturns = struct {
		result1 cf.InstanceState
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetInstanceStateReturnsOnCall(i int, result1 cf.InstanceState, result2 error) {
	fake.getInstanceStateMutex.Lock()
	defer fake.getInstanceStateMutex.Unlock()
	fake.GetInstanceStateStub = nil
	if fake.getInstanceStateReturnsOnCall == nil {
		fake.getInstanceStateReturnsOnCall = make(map[int]struct {
			result1 cf.InstanceState
			result2 error
		})
	}
	fake.getInstanceStateReturnsOnCall[i] = struct {
		result1 cf.InstanceState
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOffering(arg1 string, arg2 *log.Logger) ([]service.Instance, error) {
	fake.getInstancesOfServiceOfferingMutex.Lock()
	ret, specificReturn := fake.getInstancesOfServiceOfferingReturnsOnCall[len(fake.getInstancesOfServiceOfferingArgsForCall)]
	fake.getInstancesOfServiceOfferingArgsForCall = append(fake.getInstancesOfServiceOfferingArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetInstancesOfServiceOffering", []interface{}{arg1, arg2})
	fake.getInstancesOfServiceOfferingMutex.Unlock()
	if fake.GetInstancesOfServiceOfferingStub != nil {
		return fake.GetInstancesOfServiceOfferingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstancesOfServiceOfferingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingCallCount() int {
	fake.getInstancesOfServiceOfferingMutex.RLock()
	defer fake.getInstancesOfServiceOfferingMutex.RUnlock()
	return len(fake.getInstancesOfServiceOfferingArgsForCall)
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingCalls(stub func(string, *log.Logger) ([]service.Instance, error)) {
	fake.getInstancesOfServiceOfferingMutex.Lock()
	defer fake.getInstancesOfServiceOfferingMutex.Unlock()
	fake.GetInstancesOfServiceOfferingStub = stub
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingArgsForCall(i int) (string, *log.Logger) {
	fake.getInstancesOfServiceOfferingMutex.RLock()
	defer fake.getInstancesOfServiceOfferingMutex.RUnlock()
	argsForCall := fake.getInstancesOfServiceOfferingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingReturns(result1 []service.Instance, result2 error) {
	fake.getInstancesOfServiceOfferingMutex.Lock()
	defer fake.getInstancesOfServiceOfferingMutex.Unlock()
	fake.GetInstancesOfServiceOfferingStub = nil
	fake.getInstancesOfServiceOfferingReturns = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingReturnsOnCall(i int, result1 []service.Instance, result2 error) {
	fake.getInstancesOfServiceOfferingMutex.Lock()
	defer fake.getInstancesOfServiceOfferingMutex.Unlock()
	fake.GetInstancesOfServiceOfferingStub = nil
	if fake.getInstancesOfServiceOfferingReturnsOnCall == nil {
		fake.getInstancesOfServiceOfferingReturnsOnCall = make(map[int]struct {
			result1 []service.Instance
			result2 error
		})
	}
	fake.getInstancesOfServiceOfferingReturnsOnCall[i] = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingByOrgSpace(arg1 string, arg2 string, arg3 string, arg4 *log.Logger) ([]service.Instance, error) {
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Lock()
	ret, specificReturn := fake.getInstancesOfServiceOfferingByOrgSpaceReturnsOnCall[len(fake.getInstancesOfServiceOfferingByOrgSpaceArgsForCall)]
	fake.getInstancesOfServiceOfferingByOrgSpaceArgsForCall = append(fake.getInstancesOfServiceOfferingByOrgSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 *log.Logger
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetInstancesOfServiceOfferingByOrgSpace", []interface{}{arg1, arg2, arg3, arg4})
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Unlock()
	if fake.GetInstancesOfServiceOfferingByOrgSpaceStub != nil {
		return fake.GetInstancesOfServiceOfferingByOrgSpaceStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstancesOfServiceOfferingByOrgSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingByOrgSpaceCallCount() int {
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.RLock()
	defer fake.getInstancesOfServiceOfferingByOrgSpaceMutex.RUnlock()
	return len(fake.getInstancesOfServiceOfferingByOrgSpaceArgsForCall)
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingByOrgSpaceCalls(stub func(string, string, string, *log.Logger) ([]service.Instance, error)) {
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Lock()
	defer fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Unlock()
	fake.GetInstancesOfServiceOfferingByOrgSpaceStub = stub
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingByOrgSpaceArgsForCall(i int) (string, string, string, *log.Logger) {
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.RLock()
	defer fake.getInstancesOfServiceOfferingByOrgSpaceMutex.RUnlock()
	argsForCall := fake.getInstancesOfServiceOfferingByOrgSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingByOrgSpaceReturns(result1 []service.Instance, result2 error) {
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Lock()
	defer fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Unlock()
	fake.GetInstancesOfServiceOfferingByOrgSpaceStub = nil
	fake.getInstancesOfServiceOfferingByOrgSpaceReturns = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetInstancesOfServiceOfferingByOrgSpaceReturnsOnCall(i int, result1 []service.Instance, result2 error) {
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Lock()
	defer fake.getInstancesOfServiceOfferingByOrgSpaceMutex.Unlock()
	fake.GetInstancesOfServiceOfferingByOrgSpaceStub = nil
	if fake.getInstancesOfServiceOfferingByOrgSpaceReturnsOnCall == nil {
		fake.getInstancesOfServiceOfferingByOrgSpaceReturnsOnCall = make(map[int]struct {
			result1 []service.Instance
			result2 error
		})
	}
	fake.getInstancesOfServiceOfferingByOrgSpaceReturnsOnCall[i] = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.countInstancesOfPlanMutex.RLock()
	defer fake.countInstancesOfPlanMutex.RUnlock()
	fake.countInstancesOfServiceOfferingMutex.RLock()
	defer fake.countInstancesOfServiceOfferingMutex.RUnlock()
	fake.getAPIVersionMutex.RLock()
	defer fake.getAPIVersionMutex.RUnlock()
	fake.getInstanceStateMutex.RLock()
	defer fake.getInstanceStateMutex.RUnlock()
	fake.getInstancesOfServiceOfferingMutex.RLock()
	defer fake.getInstancesOfServiceOfferingMutex.RUnlock()
	fake.getInstancesOfServiceOfferingByOrgSpaceMutex.RLock()
	defer fake.getInstancesOfServiceOfferingByOrgSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCloudFoundryClient) recordInvocation(key string, args []interface{}) {
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

var _ broker.CloudFoundryClient = new(FakeCloudFoundryClient)
