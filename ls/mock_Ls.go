// Code generated by mockery v2.51.1. DO NOT EDIT.

package ls

import mock "github.com/stretchr/testify/mock"

// MockLs is an autogenerated mock type for the Ls type
type MockLs struct {
	mock.Mock
}

type MockLs_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLs) EXPECT() *MockLs_Expecter {
	return &MockLs_Expecter{mock: &_m.Mock}
}

// ListDirectory provides a mock function with given fields: name
func (_m *MockLs) ListDirectory(name string) (string, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for ListDirectory")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLs_ListDirectory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDirectory'
type MockLs_ListDirectory_Call struct {
	*mock.Call
}

// ListDirectory is a helper method to define mock.On call
//   - name string
func (_e *MockLs_Expecter) ListDirectory(name interface{}) *MockLs_ListDirectory_Call {
	return &MockLs_ListDirectory_Call{Call: _e.mock.On("ListDirectory", name)}
}

func (_c *MockLs_ListDirectory_Call) Run(run func(name string)) *MockLs_ListDirectory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLs_ListDirectory_Call) Return(_a0 string, _a1 error) *MockLs_ListDirectory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLs_ListDirectory_Call) RunAndReturn(run func(string) (string, error)) *MockLs_ListDirectory_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLs creates a new instance of MockLs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLs(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLs {
	mock := &MockLs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
