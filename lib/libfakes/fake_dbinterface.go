// Code generated by counterfeiter. DO NOT EDIT.
package libfakes

import (
	"sync"

	"github.com/petewall/device-service/v2/lib"
)

type FakeDBInterface struct {
	GetDeviceStub        func(string) (*lib.Device, error)
	getDeviceMutex       sync.RWMutex
	getDeviceArgsForCall []struct {
		arg1 string
	}
	getDeviceReturns struct {
		result1 *lib.Device
		result2 error
	}
	getDeviceReturnsOnCall map[int]struct {
		result1 *lib.Device
		result2 error
	}
	GetDevicesStub        func() ([]*lib.Device, error)
	getDevicesMutex       sync.RWMutex
	getDevicesArgsForCall []struct {
	}
	getDevicesReturns struct {
		result1 []*lib.Device
		result2 error
	}
	getDevicesReturnsOnCall map[int]struct {
		result1 []*lib.Device
		result2 error
	}
	UpdateDeviceStub        func(*lib.Device) error
	updateDeviceMutex       sync.RWMutex
	updateDeviceArgsForCall []struct {
		arg1 *lib.Device
	}
	updateDeviceReturns struct {
		result1 error
	}
	updateDeviceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBInterface) GetDevice(arg1 string) (*lib.Device, error) {
	fake.getDeviceMutex.Lock()
	ret, specificReturn := fake.getDeviceReturnsOnCall[len(fake.getDeviceArgsForCall)]
	fake.getDeviceArgsForCall = append(fake.getDeviceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetDeviceStub
	fakeReturns := fake.getDeviceReturns
	fake.recordInvocation("GetDevice", []interface{}{arg1})
	fake.getDeviceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDBInterface) GetDeviceCallCount() int {
	fake.getDeviceMutex.RLock()
	defer fake.getDeviceMutex.RUnlock()
	return len(fake.getDeviceArgsForCall)
}

func (fake *FakeDBInterface) GetDeviceCalls(stub func(string) (*lib.Device, error)) {
	fake.getDeviceMutex.Lock()
	defer fake.getDeviceMutex.Unlock()
	fake.GetDeviceStub = stub
}

func (fake *FakeDBInterface) GetDeviceArgsForCall(i int) string {
	fake.getDeviceMutex.RLock()
	defer fake.getDeviceMutex.RUnlock()
	argsForCall := fake.getDeviceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDBInterface) GetDeviceReturns(result1 *lib.Device, result2 error) {
	fake.getDeviceMutex.Lock()
	defer fake.getDeviceMutex.Unlock()
	fake.GetDeviceStub = nil
	fake.getDeviceReturns = struct {
		result1 *lib.Device
		result2 error
	}{result1, result2}
}

func (fake *FakeDBInterface) GetDeviceReturnsOnCall(i int, result1 *lib.Device, result2 error) {
	fake.getDeviceMutex.Lock()
	defer fake.getDeviceMutex.Unlock()
	fake.GetDeviceStub = nil
	if fake.getDeviceReturnsOnCall == nil {
		fake.getDeviceReturnsOnCall = make(map[int]struct {
			result1 *lib.Device
			result2 error
		})
	}
	fake.getDeviceReturnsOnCall[i] = struct {
		result1 *lib.Device
		result2 error
	}{result1, result2}
}

func (fake *FakeDBInterface) GetDevices() ([]*lib.Device, error) {
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

func (fake *FakeDBInterface) GetDevicesCalls(stub func() ([]*lib.Device, error)) {
	fake.getDevicesMutex.Lock()
	defer fake.getDevicesMutex.Unlock()
	fake.GetDevicesStub = stub
}

func (fake *FakeDBInterface) GetDevicesReturns(result1 []*lib.Device, result2 error) {
	fake.getDevicesMutex.Lock()
	defer fake.getDevicesMutex.Unlock()
	fake.GetDevicesStub = nil
	fake.getDevicesReturns = struct {
		result1 []*lib.Device
		result2 error
	}{result1, result2}
}

func (fake *FakeDBInterface) GetDevicesReturnsOnCall(i int, result1 []*lib.Device, result2 error) {
	fake.getDevicesMutex.Lock()
	defer fake.getDevicesMutex.Unlock()
	fake.GetDevicesStub = nil
	if fake.getDevicesReturnsOnCall == nil {
		fake.getDevicesReturnsOnCall = make(map[int]struct {
			result1 []*lib.Device
			result2 error
		})
	}
	fake.getDevicesReturnsOnCall[i] = struct {
		result1 []*lib.Device
		result2 error
	}{result1, result2}
}

func (fake *FakeDBInterface) UpdateDevice(arg1 *lib.Device) error {
	fake.updateDeviceMutex.Lock()
	ret, specificReturn := fake.updateDeviceReturnsOnCall[len(fake.updateDeviceArgsForCall)]
	fake.updateDeviceArgsForCall = append(fake.updateDeviceArgsForCall, struct {
		arg1 *lib.Device
	}{arg1})
	stub := fake.UpdateDeviceStub
	fakeReturns := fake.updateDeviceReturns
	fake.recordInvocation("UpdateDevice", []interface{}{arg1})
	fake.updateDeviceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDBInterface) UpdateDeviceCallCount() int {
	fake.updateDeviceMutex.RLock()
	defer fake.updateDeviceMutex.RUnlock()
	return len(fake.updateDeviceArgsForCall)
}

func (fake *FakeDBInterface) UpdateDeviceCalls(stub func(*lib.Device) error) {
	fake.updateDeviceMutex.Lock()
	defer fake.updateDeviceMutex.Unlock()
	fake.UpdateDeviceStub = stub
}

func (fake *FakeDBInterface) UpdateDeviceArgsForCall(i int) *lib.Device {
	fake.updateDeviceMutex.RLock()
	defer fake.updateDeviceMutex.RUnlock()
	argsForCall := fake.updateDeviceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDBInterface) UpdateDeviceReturns(result1 error) {
	fake.updateDeviceMutex.Lock()
	defer fake.updateDeviceMutex.Unlock()
	fake.UpdateDeviceStub = nil
	fake.updateDeviceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBInterface) UpdateDeviceReturnsOnCall(i int, result1 error) {
	fake.updateDeviceMutex.Lock()
	defer fake.updateDeviceMutex.Unlock()
	fake.UpdateDeviceStub = nil
	if fake.updateDeviceReturnsOnCall == nil {
		fake.updateDeviceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateDeviceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDeviceMutex.RLock()
	defer fake.getDeviceMutex.RUnlock()
	fake.getDevicesMutex.RLock()
	defer fake.getDevicesMutex.RUnlock()
	fake.updateDeviceMutex.RLock()
	defer fake.updateDeviceMutex.RUnlock()
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

var _ lib.DBInterface = new(FakeDBInterface)
