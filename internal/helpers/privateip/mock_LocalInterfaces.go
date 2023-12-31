// Code generated by mockery v2.36.0. DO NOT EDIT.

package privateip

import mock "github.com/stretchr/testify/mock"

// MockLocalInterfaces is an autogenerated mock type for the LocalInterfaces type
type MockLocalInterfaces struct {
	mock.Mock
}

type MockLocalInterfaces_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLocalInterfaces) EXPECT() *MockLocalInterfaces_Expecter {
	return &MockLocalInterfaces_Expecter{mock: &_m.Mock}
}

// GetFirst provides a mock function with given fields:
func (_m *MockLocalInterfaces) GetFirst() *NIC {
	ret := _m.Called()

	var r0 *NIC
	if rf, ok := ret.Get(0).(func() *NIC); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*NIC)
		}
	}

	return r0
}

// MockLocalInterfaces_GetFirst_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirst'
type MockLocalInterfaces_GetFirst_Call struct {
	*mock.Call
}

// GetFirst is a helper method to define mock.On call
func (_e *MockLocalInterfaces_Expecter) GetFirst() *MockLocalInterfaces_GetFirst_Call {
	return &MockLocalInterfaces_GetFirst_Call{Call: _e.mock.On("GetFirst")}
}

func (_c *MockLocalInterfaces_GetFirst_Call) Run(run func()) *MockLocalInterfaces_GetFirst_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLocalInterfaces_GetFirst_Call) Return(_a0 *NIC) *MockLocalInterfaces_GetFirst_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLocalInterfaces_GetFirst_Call) RunAndReturn(run func() *NIC) *MockLocalInterfaces_GetFirst_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrimary provides a mock function with given fields:
func (_m *MockLocalInterfaces) GetPrimary() *NIC {
	ret := _m.Called()

	var r0 *NIC
	if rf, ok := ret.Get(0).(func() *NIC); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*NIC)
		}
	}

	return r0
}

// MockLocalInterfaces_GetPrimary_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrimary'
type MockLocalInterfaces_GetPrimary_Call struct {
	*mock.Call
}

// GetPrimary is a helper method to define mock.On call
func (_e *MockLocalInterfaces_Expecter) GetPrimary() *MockLocalInterfaces_GetPrimary_Call {
	return &MockLocalInterfaces_GetPrimary_Call{Call: _e.mock.On("GetPrimary")}
}

func (_c *MockLocalInterfaces_GetPrimary_Call) Run(run func()) *MockLocalInterfaces_GetPrimary_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLocalInterfaces_GetPrimary_Call) Return(_a0 *NIC) *MockLocalInterfaces_GetPrimary_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLocalInterfaces_GetPrimary_Call) RunAndReturn(run func() *NIC) *MockLocalInterfaces_GetPrimary_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrimaryAbsentReason provides a mock function with given fields:
func (_m *MockLocalInterfaces) GetPrimaryAbsentReason() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockLocalInterfaces_GetPrimaryAbsentReason_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrimaryAbsentReason'
type MockLocalInterfaces_GetPrimaryAbsentReason_Call struct {
	*mock.Call
}

// GetPrimaryAbsentReason is a helper method to define mock.On call
func (_e *MockLocalInterfaces_Expecter) GetPrimaryAbsentReason() *MockLocalInterfaces_GetPrimaryAbsentReason_Call {
	return &MockLocalInterfaces_GetPrimaryAbsentReason_Call{Call: _e.mock.On("GetPrimaryAbsentReason")}
}

func (_c *MockLocalInterfaces_GetPrimaryAbsentReason_Call) Run(run func()) *MockLocalInterfaces_GetPrimaryAbsentReason_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLocalInterfaces_GetPrimaryAbsentReason_Call) Return(_a0 string) *MockLocalInterfaces_GetPrimaryAbsentReason_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLocalInterfaces_GetPrimaryAbsentReason_Call) RunAndReturn(run func() string) *MockLocalInterfaces_GetPrimaryAbsentReason_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecondaries provides a mock function with given fields:
func (_m *MockLocalInterfaces) GetSecondaries() []*NIC {
	ret := _m.Called()

	var r0 []*NIC
	if rf, ok := ret.Get(0).(func() []*NIC); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*NIC)
		}
	}

	return r0
}

// MockLocalInterfaces_GetSecondaries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecondaries'
type MockLocalInterfaces_GetSecondaries_Call struct {
	*mock.Call
}

// GetSecondaries is a helper method to define mock.On call
func (_e *MockLocalInterfaces_Expecter) GetSecondaries() *MockLocalInterfaces_GetSecondaries_Call {
	return &MockLocalInterfaces_GetSecondaries_Call{Call: _e.mock.On("GetSecondaries")}
}

func (_c *MockLocalInterfaces_GetSecondaries_Call) Run(run func()) *MockLocalInterfaces_GetSecondaries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLocalInterfaces_GetSecondaries_Call) Return(_a0 []*NIC) *MockLocalInterfaces_GetSecondaries_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLocalInterfaces_GetSecondaries_Call) RunAndReturn(run func() []*NIC) *MockLocalInterfaces_GetSecondaries_Call {
	_c.Call.Return(run)
	return _c
}

// ScanInterfaces provides a mock function with given fields:
func (_m *MockLocalInterfaces) ScanInterfaces() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLocalInterfaces_ScanInterfaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScanInterfaces'
type MockLocalInterfaces_ScanInterfaces_Call struct {
	*mock.Call
}

// ScanInterfaces is a helper method to define mock.On call
func (_e *MockLocalInterfaces_Expecter) ScanInterfaces() *MockLocalInterfaces_ScanInterfaces_Call {
	return &MockLocalInterfaces_ScanInterfaces_Call{Call: _e.mock.On("ScanInterfaces")}
}

func (_c *MockLocalInterfaces_ScanInterfaces_Call) Run(run func()) *MockLocalInterfaces_ScanInterfaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLocalInterfaces_ScanInterfaces_Call) Return(_a0 error) *MockLocalInterfaces_ScanInterfaces_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLocalInterfaces_ScanInterfaces_Call) RunAndReturn(run func() error) *MockLocalInterfaces_ScanInterfaces_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLocalInterfaces creates a new instance of MockLocalInterfaces. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLocalInterfaces(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLocalInterfaces {
	mock := &MockLocalInterfaces{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
