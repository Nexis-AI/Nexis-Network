// Code generated by mockery. DO NOT EDIT.

package etherman

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// ethGasStationMock is an autogenerated mock type for the GasPricer type
type ethGasStationMock struct {
	mock.Mock
}

type ethGasStationMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ethGasStationMock) EXPECT() *ethGasStationMock_Expecter {
	return &ethGasStationMock_Expecter{mock: &_m.Mock}
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *ethGasStationMock) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuggestGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethGasStationMock_SuggestGasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestGasPrice'
type ethGasStationMock_SuggestGasPrice_Call struct {
	*mock.Call
}

// SuggestGasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ethGasStationMock_Expecter) SuggestGasPrice(ctx interface{}) *ethGasStationMock_SuggestGasPrice_Call {
	return &ethGasStationMock_SuggestGasPrice_Call{Call: _e.mock.On("SuggestGasPrice", ctx)}
}

func (_c *ethGasStationMock_SuggestGasPrice_Call) Run(run func(ctx context.Context)) *ethGasStationMock_SuggestGasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ethGasStationMock_SuggestGasPrice_Call) Return(_a0 *big.Int, _a1 error) *ethGasStationMock_SuggestGasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethGasStationMock_SuggestGasPrice_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *ethGasStationMock_SuggestGasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// newEthGasStationMock creates a new instance of ethGasStationMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newEthGasStationMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ethGasStationMock {
	mock := &ethGasStationMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
