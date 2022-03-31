// Code generated by counterfeiter. DO NOT EDIT.
package deviceservicefakes

import (
	"sync"

	deviceservice "github.com/petewall/device-service/v2"
)

type FakeDBInterface struct {
	GetDevicesStub        func() (error, []*deviceservice.Device)
	getDevicesMutex       sync.RWMutex
	getDevicesArgsForCall []struct {
	}
	getDevicesReturns struct {
		result1 error
		result2 []*deviceservice.Device
	}
	getDevicesReturnsOnCall map[int]struct {
		result1 error
		result2 []*deviceservice.Device
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBInterface) GetDevices() (error, []*deviceservice.Device) {
	fake.getDevicesMutex.Lock()
	ret, specificReturn := fake.getDevicesReturnsOnCall[len(fake.getDevicesArgsForCall)]
	fake.getDevicesArgsForCall = append(fake.getDevicesArgsForCall, struct {
	}{})
	stub := fake.GetDevicesStub
	fakeReturns := fake.getDevicesReturns
	fake.recordInvocation("GetDevices", []interface{}{})
	fake.getDevicesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDBInterface) GetDevicesCallCount() int {
	fake.getDevicesMutex.RLock()
	defer fake.getDevicesMutex.RUnlock()
	return len(fake.getDevicesArgsForCall)
}

func (fake *FakeDBInterface) GetDevicesCalls(stub func() (error, []*deviceservice.Device)) {
	fake.getDevicesMutex.Lock()
	defer fake.getDevicesMutex.Unlock()
	fake.GetDevicesStub = stub
}

func (fake *FakeDBInterface) GetDevicesReturns(result1 error, result2 []*deviceservice.Device) {
	fake.getDevicesMutex.Lock()
	defer fake.getDevicesMutex.Unlock()
	fake.GetDevicesStub = nil
	fake.getDevicesReturns = struct {
		result1 error
		result2 []*deviceservice.Device
	}{result1, result2}
}

func (fake *FakeDBInterface) GetDevicesReturnsOnCall(i int, result1 error, result2 []*deviceservice.Device) {
	fake.getDevicesMutex.Lock()
	defer fake.getDevicesMutex.Unlock()
	fake.GetDevicesStub = nil
	if fake.getDevicesReturnsOnCall == nil {
		fake.getDevicesReturnsOnCall = make(map[int]struct {
			result1 error
			result2 []*deviceservice.Device
		})
	}
	fake.getDevicesReturnsOnCall[i] = struct {
		result1 error
		result2 []*deviceservice.Device
	}{result1, result2}
}

func (fake *FakeDBInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDevicesMutex.RLock()
	defer fake.getDevicesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDBInterface) recordInvocation(key string, args []interface{}) {
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

var _ deviceservice.DBInterface = new(FakeDBInterface)
