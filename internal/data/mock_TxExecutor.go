// Code generated by mockery v2.51.1. DO NOT EDIT.

package data

import mock "github.com/stretchr/testify/mock"

// MockTxExecutor is an autogenerated mock type for the TxExecutor type
type MockTxExecutor[T interface{}] struct {
	mock.Mock
}

type MockTxExecutor_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockTxExecutor[T]) EXPECT() *MockTxExecutor_Expecter[T] {
	return &MockTxExecutor_Expecter[T]{mock: &_m.Mock}
}

// Commit provides a mock function with no fields
func (_m *MockTxExecutor[T]) Commit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTxExecutor_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type MockTxExecutor_Commit_Call[T interface{}] struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
func (_e *MockTxExecutor_Expecter[T]) Commit() *MockTxExecutor_Commit_Call[T] {
	return &MockTxExecutor_Commit_Call[T]{Call: _e.mock.On("Commit")}
}

func (_c *MockTxExecutor_Commit_Call[T]) Run(run func()) *MockTxExecutor_Commit_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTxExecutor_Commit_Call[T]) Return(_a0 error) *MockTxExecutor_Commit_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTxExecutor_Commit_Call[T]) RunAndReturn(run func() error) *MockTxExecutor_Commit_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Executor provides a mock function with no fields
func (_m *MockTxExecutor[T]) Executor() T {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Executor")
	}

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(T)
		}
	}

	return r0
}

// MockTxExecutor_Executor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Executor'
type MockTxExecutor_Executor_Call[T interface{}] struct {
	*mock.Call
}

// Executor is a helper method to define mock.On call
func (_e *MockTxExecutor_Expecter[T]) Executor() *MockTxExecutor_Executor_Call[T] {
	return &MockTxExecutor_Executor_Call[T]{Call: _e.mock.On("Executor")}
}

func (_c *MockTxExecutor_Executor_Call[T]) Run(run func()) *MockTxExecutor_Executor_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTxExecutor_Executor_Call[T]) Return(_a0 T) *MockTxExecutor_Executor_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTxExecutor_Executor_Call[T]) RunAndReturn(run func() T) *MockTxExecutor_Executor_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with no fields
func (_m *MockTxExecutor[T]) Rollback() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTxExecutor_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type MockTxExecutor_Rollback_Call[T interface{}] struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
func (_e *MockTxExecutor_Expecter[T]) Rollback() *MockTxExecutor_Rollback_Call[T] {
	return &MockTxExecutor_Rollback_Call[T]{Call: _e.mock.On("Rollback")}
}

func (_c *MockTxExecutor_Rollback_Call[T]) Run(run func()) *MockTxExecutor_Rollback_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTxExecutor_Rollback_Call[T]) Return(_a0 error) *MockTxExecutor_Rollback_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTxExecutor_Rollback_Call[T]) RunAndReturn(run func() error) *MockTxExecutor_Rollback_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockTxExecutor creates a new instance of MockTxExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTxExecutor[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTxExecutor[T] {
	mock := &MockTxExecutor[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
