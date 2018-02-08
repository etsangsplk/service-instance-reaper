// Code generated by counterfeiter. DO NOT EDIT.
package cloudfoundryfakes

import (
	"sync"

	"github.com/pivotal-cf/service-instance-reaper/cloudfoundry"
)

type FakeClient struct {
	GetServicesStub        func(serviceName string) ([]cloudfoundry.Service, error)
	getServicesMutex       sync.RWMutex
	getServicesArgsForCall []struct {
		serviceName string
	}
	getServicesReturns struct {
		result1 []cloudfoundry.Service
		result2 error
	}
	getServicesReturnsOnCall map[int]struct {
		result1 []cloudfoundry.Service
		result2 error
	}
	GetServicePlansStub        func(serviceGuid string) ([]cloudfoundry.ServicePlan, error)
	getServicePlansMutex       sync.RWMutex
	getServicePlansArgsForCall []struct {
		serviceGuid string
	}
	getServicePlansReturns struct {
		result1 []cloudfoundry.ServicePlan
		result2 error
	}
	getServicePlansReturnsOnCall map[int]struct {
		result1 []cloudfoundry.ServicePlan
		result2 error
	}
	GetServicePlanInstancesStub        func(servicePlanGuid string) (chan cloudfoundry.ServiceInstance, chan error)
	getServicePlanInstancesMutex       sync.RWMutex
	getServicePlanInstancesArgsForCall []struct {
		servicePlanGuid string
	}
	getServicePlanInstancesReturns struct {
		result1 chan cloudfoundry.ServiceInstance
		result2 chan error
	}
	getServicePlanInstancesReturnsOnCall map[int]struct {
		result1 chan cloudfoundry.ServiceInstance
		result2 chan error
	}
	DeleteServiceInstanceStub        func(serviceInstanceGuid string, recursive bool) error
	deleteServiceInstanceMutex       sync.RWMutex
	deleteServiceInstanceArgsForCall []struct {
		serviceInstanceGuid string
		recursive           bool
	}
	deleteServiceInstanceReturns struct {
		result1 error
	}
	deleteServiceInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) GetServices(serviceName string) ([]cloudfoundry.Service, error) {
	fake.getServicesMutex.Lock()
	ret, specificReturn := fake.getServicesReturnsOnCall[len(fake.getServicesArgsForCall)]
	fake.getServicesArgsForCall = append(fake.getServicesArgsForCall, struct {
		serviceName string
	}{serviceName})
	fake.recordInvocation("GetServices", []interface{}{serviceName})
	fake.getServicesMutex.Unlock()
	if fake.GetServicesStub != nil {
		return fake.GetServicesStub(serviceName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServicesReturns.result1, fake.getServicesReturns.result2
}

func (fake *FakeClient) GetServicesCallCount() int {
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	return len(fake.getServicesArgsForCall)
}

func (fake *FakeClient) GetServicesArgsForCall(i int) string {
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	return fake.getServicesArgsForCall[i].serviceName
}

func (fake *FakeClient) GetServicesReturns(result1 []cloudfoundry.Service, result2 error) {
	fake.GetServicesStub = nil
	fake.getServicesReturns = struct {
		result1 []cloudfoundry.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetServicesReturnsOnCall(i int, result1 []cloudfoundry.Service, result2 error) {
	fake.GetServicesStub = nil
	if fake.getServicesReturnsOnCall == nil {
		fake.getServicesReturnsOnCall = make(map[int]struct {
			result1 []cloudfoundry.Service
			result2 error
		})
	}
	fake.getServicesReturnsOnCall[i] = struct {
		result1 []cloudfoundry.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetServicePlans(serviceGuid string) ([]cloudfoundry.ServicePlan, error) {
	fake.getServicePlansMutex.Lock()
	ret, specificReturn := fake.getServicePlansReturnsOnCall[len(fake.getServicePlansArgsForCall)]
	fake.getServicePlansArgsForCall = append(fake.getServicePlansArgsForCall, struct {
		serviceGuid string
	}{serviceGuid})
	fake.recordInvocation("GetServicePlans", []interface{}{serviceGuid})
	fake.getServicePlansMutex.Unlock()
	if fake.GetServicePlansStub != nil {
		return fake.GetServicePlansStub(serviceGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServicePlansReturns.result1, fake.getServicePlansReturns.result2
}

func (fake *FakeClient) GetServicePlansCallCount() int {
	fake.getServicePlansMutex.RLock()
	defer fake.getServicePlansMutex.RUnlock()
	return len(fake.getServicePlansArgsForCall)
}

func (fake *FakeClient) GetServicePlansArgsForCall(i int) string {
	fake.getServicePlansMutex.RLock()
	defer fake.getServicePlansMutex.RUnlock()
	return fake.getServicePlansArgsForCall[i].serviceGuid
}

func (fake *FakeClient) GetServicePlansReturns(result1 []cloudfoundry.ServicePlan, result2 error) {
	fake.GetServicePlansStub = nil
	fake.getServicePlansReturns = struct {
		result1 []cloudfoundry.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetServicePlansReturnsOnCall(i int, result1 []cloudfoundry.ServicePlan, result2 error) {
	fake.GetServicePlansStub = nil
	if fake.getServicePlansReturnsOnCall == nil {
		fake.getServicePlansReturnsOnCall = make(map[int]struct {
			result1 []cloudfoundry.ServicePlan
			result2 error
		})
	}
	fake.getServicePlansReturnsOnCall[i] = struct {
		result1 []cloudfoundry.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetServicePlanInstances(servicePlanGuid string) (chan cloudfoundry.ServiceInstance, chan error) {
	fake.getServicePlanInstancesMutex.Lock()
	ret, specificReturn := fake.getServicePlanInstancesReturnsOnCall[len(fake.getServicePlanInstancesArgsForCall)]
	fake.getServicePlanInstancesArgsForCall = append(fake.getServicePlanInstancesArgsForCall, struct {
		servicePlanGuid string
	}{servicePlanGuid})
	fake.recordInvocation("GetServicePlanInstances", []interface{}{servicePlanGuid})
	fake.getServicePlanInstancesMutex.Unlock()
	if fake.GetServicePlanInstancesStub != nil {
		return fake.GetServicePlanInstancesStub(servicePlanGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServicePlanInstancesReturns.result1, fake.getServicePlanInstancesReturns.result2
}

func (fake *FakeClient) GetServicePlanInstancesCallCount() int {
	fake.getServicePlanInstancesMutex.RLock()
	defer fake.getServicePlanInstancesMutex.RUnlock()
	return len(fake.getServicePlanInstancesArgsForCall)
}

func (fake *FakeClient) GetServicePlanInstancesArgsForCall(i int) string {
	fake.getServicePlanInstancesMutex.RLock()
	defer fake.getServicePlanInstancesMutex.RUnlock()
	return fake.getServicePlanInstancesArgsForCall[i].servicePlanGuid
}

func (fake *FakeClient) GetServicePlanInstancesReturns(result1 chan cloudfoundry.ServiceInstance, result2 chan error) {
	fake.GetServicePlanInstancesStub = nil
	fake.getServicePlanInstancesReturns = struct {
		result1 chan cloudfoundry.ServiceInstance
		result2 chan error
	}{result1, result2}
}

func (fake *FakeClient) GetServicePlanInstancesReturnsOnCall(i int, result1 chan cloudfoundry.ServiceInstance, result2 chan error) {
	fake.GetServicePlanInstancesStub = nil
	if fake.getServicePlanInstancesReturnsOnCall == nil {
		fake.getServicePlanInstancesReturnsOnCall = make(map[int]struct {
			result1 chan cloudfoundry.ServiceInstance
			result2 chan error
		})
	}
	fake.getServicePlanInstancesReturnsOnCall[i] = struct {
		result1 chan cloudfoundry.ServiceInstance
		result2 chan error
	}{result1, result2}
}

func (fake *FakeClient) DeleteServiceInstance(serviceInstanceGuid string, recursive bool) error {
	fake.deleteServiceInstanceMutex.Lock()
	ret, specificReturn := fake.deleteServiceInstanceReturnsOnCall[len(fake.deleteServiceInstanceArgsForCall)]
	fake.deleteServiceInstanceArgsForCall = append(fake.deleteServiceInstanceArgsForCall, struct {
		serviceInstanceGuid string
		recursive           bool
	}{serviceInstanceGuid, recursive})
	fake.recordInvocation("DeleteServiceInstance", []interface{}{serviceInstanceGuid, recursive})
	fake.deleteServiceInstanceMutex.Unlock()
	if fake.DeleteServiceInstanceStub != nil {
		return fake.DeleteServiceInstanceStub(serviceInstanceGuid, recursive)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteServiceInstanceReturns.result1
}

func (fake *FakeClient) DeleteServiceInstanceCallCount() int {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	return len(fake.deleteServiceInstanceArgsForCall)
}

func (fake *FakeClient) DeleteServiceInstanceArgsForCall(i int) (string, bool) {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	return fake.deleteServiceInstanceArgsForCall[i].serviceInstanceGuid, fake.deleteServiceInstanceArgsForCall[i].recursive
}

func (fake *FakeClient) DeleteServiceInstanceReturns(result1 error) {
	fake.DeleteServiceInstanceStub = nil
	fake.deleteServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteServiceInstanceReturnsOnCall(i int, result1 error) {
	fake.DeleteServiceInstanceStub = nil
	if fake.deleteServiceInstanceReturnsOnCall == nil {
		fake.deleteServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteServiceInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	fake.getServicePlansMutex.RLock()
	defer fake.getServicePlansMutex.RUnlock()
	fake.getServicePlanInstancesMutex.RLock()
	defer fake.getServicePlanInstancesMutex.RUnlock()
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ cloudfoundry.Client = new(FakeClient)
