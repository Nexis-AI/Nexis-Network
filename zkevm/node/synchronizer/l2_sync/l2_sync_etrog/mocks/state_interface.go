// Code generated by mockery. DO NOT EDIT.

package mock_l2_sync_etrog

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/0xPolygonHermez/zkevm-node/state"
)

// StateInterface is an autogenerated mock type for the StateInterface type
type StateInterface struct {
	mock.Mock
}

type StateInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StateInterface) EXPECT() *StateInterface_Expecter {
	return &StateInterface_Expecter{mock: &_m.Mock}
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateInterface) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginStateTransaction")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInterface_BeginStateTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginStateTransaction'
type StateInterface_BeginStateTransaction_Call struct {
	*mock.Call
}

// BeginStateTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *StateInterface_Expecter) BeginStateTransaction(ctx interface{}) *StateInterface_BeginStateTransaction_Call {
	return &StateInterface_BeginStateTransaction_Call{Call: _e.mock.On("BeginStateTransaction", ctx)}
}

func (_c *StateInterface_BeginStateTransaction_Call) Run(run func(ctx context.Context)) *StateInterface_BeginStateTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *StateInterface_BeginStateTransaction_Call) Return(_a0 pgx.Tx, _a1 error) *StateInterface_BeginStateTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateInterface_BeginStateTransaction_Call) RunAndReturn(run func(context.Context) (pgx.Tx, error)) *StateInterface_BeginStateTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// CloseBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *StateInterface) CloseBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for CloseBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateInterface_CloseBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseBatch'
type StateInterface_CloseBatch_Call struct {
	*mock.Call
}

// CloseBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - receipt state.ProcessingReceipt
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) CloseBatch(ctx interface{}, receipt interface{}, dbTx interface{}) *StateInterface_CloseBatch_Call {
	return &StateInterface_CloseBatch_Call{Call: _e.mock.On("CloseBatch", ctx, receipt, dbTx)}
}

func (_c *StateInterface_CloseBatch_Call) Run(run func(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx)) *StateInterface_CloseBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(state.ProcessingReceipt), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_CloseBatch_Call) Return(_a0 error) *StateInterface_CloseBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateInterface_CloseBatch_Call) RunAndReturn(run func(context.Context, state.ProcessingReceipt, pgx.Tx) error) *StateInterface_CloseBatch_Call {
	_c.Call.Return(run)
	return _c
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateInterface) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetBatchByNumber")
	}

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInterface_GetBatchByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBatchByNumber'
type StateInterface_GetBatchByNumber_Call struct {
	*mock.Call
}

// GetBatchByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) GetBatchByNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *StateInterface_GetBatchByNumber_Call {
	return &StateInterface_GetBatchByNumber_Call{Call: _e.mock.On("GetBatchByNumber", ctx, batchNumber, dbTx)}
}

func (_c *StateInterface_GetBatchByNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx pgx.Tx)) *StateInterface_GetBatchByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_GetBatchByNumber_Call) Return(_a0 *state.Batch, _a1 error) *StateInterface_GetBatchByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateInterface_GetBatchByNumber_Call) RunAndReturn(run func(context.Context, uint64, pgx.Tx) (*state.Batch, error)) *StateInterface_GetBatchByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetForkIDByBatchNumber provides a mock function with given fields: batchNumber
func (_m *StateInterface) GetForkIDByBatchNumber(batchNumber uint64) uint64 {
	ret := _m.Called(batchNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetForkIDByBatchNumber")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(batchNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// StateInterface_GetForkIDByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForkIDByBatchNumber'
type StateInterface_GetForkIDByBatchNumber_Call struct {
	*mock.Call
}

// GetForkIDByBatchNumber is a helper method to define mock.On call
//   - batchNumber uint64
func (_e *StateInterface_Expecter) GetForkIDByBatchNumber(batchNumber interface{}) *StateInterface_GetForkIDByBatchNumber_Call {
	return &StateInterface_GetForkIDByBatchNumber_Call{Call: _e.mock.On("GetForkIDByBatchNumber", batchNumber)}
}

func (_c *StateInterface_GetForkIDByBatchNumber_Call) Run(run func(batchNumber uint64)) *StateInterface_GetForkIDByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *StateInterface_GetForkIDByBatchNumber_Call) Return(_a0 uint64) *StateInterface_GetForkIDByBatchNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateInterface_GetForkIDByBatchNumber_Call) RunAndReturn(run func(uint64) uint64) *StateInterface_GetForkIDByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetL1InfoTreeDataFromBatchL2Data provides a mock function with given fields: ctx, batchL2Data, dbTx
func (_m *StateInterface) GetL1InfoTreeDataFromBatchL2Data(ctx context.Context, batchL2Data []byte, dbTx pgx.Tx) (map[uint32]state.L1DataV2, common.Hash, common.Hash, error) {
	ret := _m.Called(ctx, batchL2Data, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetL1InfoTreeDataFromBatchL2Data")
	}

	var r0 map[uint32]state.L1DataV2
	var r1 common.Hash
	var r2 common.Hash
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, pgx.Tx) (map[uint32]state.L1DataV2, common.Hash, common.Hash, error)); ok {
		return rf(ctx, batchL2Data, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, pgx.Tx) map[uint32]state.L1DataV2); ok {
		r0 = rf(ctx, batchL2Data, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]state.L1DataV2)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, pgx.Tx) common.Hash); ok {
		r1 = rf(ctx, batchL2Data, dbTx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(common.Hash)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, pgx.Tx) common.Hash); ok {
		r2 = rf(ctx, batchL2Data, dbTx)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(common.Hash)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, []byte, pgx.Tx) error); ok {
		r3 = rf(ctx, batchL2Data, dbTx)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1InfoTreeDataFromBatchL2Data'
type StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call struct {
	*mock.Call
}

// GetL1InfoTreeDataFromBatchL2Data is a helper method to define mock.On call
//   - ctx context.Context
//   - batchL2Data []byte
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) GetL1InfoTreeDataFromBatchL2Data(ctx interface{}, batchL2Data interface{}, dbTx interface{}) *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call {
	return &StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call{Call: _e.mock.On("GetL1InfoTreeDataFromBatchL2Data", ctx, batchL2Data, dbTx)}
}

func (_c *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call) Run(run func(ctx context.Context, batchL2Data []byte, dbTx pgx.Tx)) *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call) Return(_a0 map[uint32]state.L1DataV2, _a1 common.Hash, _a2 common.Hash, _a3 error) *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call) RunAndReturn(run func(context.Context, []byte, pgx.Tx) (map[uint32]state.L1DataV2, common.Hash, common.Hash, error)) *StateInterface_GetL1InfoTreeDataFromBatchL2Data_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *StateInterface) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastVirtualBatchNum")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInterface_GetLastVirtualBatchNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastVirtualBatchNum'
type StateInterface_GetLastVirtualBatchNum_Call struct {
	*mock.Call
}

// GetLastVirtualBatchNum is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) GetLastVirtualBatchNum(ctx interface{}, dbTx interface{}) *StateInterface_GetLastVirtualBatchNum_Call {
	return &StateInterface_GetLastVirtualBatchNum_Call{Call: _e.mock.On("GetLastVirtualBatchNum", ctx, dbTx)}
}

func (_c *StateInterface_GetLastVirtualBatchNum_Call) Run(run func(ctx context.Context, dbTx pgx.Tx)) *StateInterface_GetLastVirtualBatchNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_GetLastVirtualBatchNum_Call) Return(_a0 uint64, _a1 error) *StateInterface_GetLastVirtualBatchNum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateInterface_GetLastVirtualBatchNum_Call) RunAndReturn(run func(context.Context, pgx.Tx) (uint64, error)) *StateInterface_GetLastVirtualBatchNum_Call {
	_c.Call.Return(run)
	return _c
}

// OpenBatch provides a mock function with given fields: ctx, processingContext, dbTx
func (_m *StateInterface) OpenBatch(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, processingContext, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for OpenBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, pgx.Tx) error); ok {
		r0 = rf(ctx, processingContext, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateInterface_OpenBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OpenBatch'
type StateInterface_OpenBatch_Call struct {
	*mock.Call
}

// OpenBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - processingContext state.ProcessingContext
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) OpenBatch(ctx interface{}, processingContext interface{}, dbTx interface{}) *StateInterface_OpenBatch_Call {
	return &StateInterface_OpenBatch_Call{Call: _e.mock.On("OpenBatch", ctx, processingContext, dbTx)}
}

func (_c *StateInterface_OpenBatch_Call) Run(run func(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx)) *StateInterface_OpenBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(state.ProcessingContext), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_OpenBatch_Call) Return(_a0 error) *StateInterface_OpenBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateInterface_OpenBatch_Call) RunAndReturn(run func(context.Context, state.ProcessingContext, pgx.Tx) error) *StateInterface_OpenBatch_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessBatchV2 provides a mock function with given fields: ctx, request, updateMerkleTree
func (_m *StateInterface) ProcessBatchV2(ctx context.Context, request state.ProcessRequest, updateMerkleTree bool) (*state.ProcessBatchResponse, string, error) {
	ret := _m.Called(ctx, request, updateMerkleTree)

	if len(ret) == 0 {
		panic("no return value specified for ProcessBatchV2")
	}

	var r0 *state.ProcessBatchResponse
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest, bool) (*state.ProcessBatchResponse, string, error)); ok {
		return rf(ctx, request, updateMerkleTree)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest, bool) *state.ProcessBatchResponse); ok {
		r0 = rf(ctx, request, updateMerkleTree)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.ProcessRequest, bool) string); ok {
		r1 = rf(ctx, request, updateMerkleTree)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, state.ProcessRequest, bool) error); ok {
		r2 = rf(ctx, request, updateMerkleTree)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StateInterface_ProcessBatchV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessBatchV2'
type StateInterface_ProcessBatchV2_Call struct {
	*mock.Call
}

// ProcessBatchV2 is a helper method to define mock.On call
//   - ctx context.Context
//   - request state.ProcessRequest
//   - updateMerkleTree bool
func (_e *StateInterface_Expecter) ProcessBatchV2(ctx interface{}, request interface{}, updateMerkleTree interface{}) *StateInterface_ProcessBatchV2_Call {
	return &StateInterface_ProcessBatchV2_Call{Call: _e.mock.On("ProcessBatchV2", ctx, request, updateMerkleTree)}
}

func (_c *StateInterface_ProcessBatchV2_Call) Run(run func(ctx context.Context, request state.ProcessRequest, updateMerkleTree bool)) *StateInterface_ProcessBatchV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(state.ProcessRequest), args[2].(bool))
	})
	return _c
}

func (_c *StateInterface_ProcessBatchV2_Call) Return(_a0 *state.ProcessBatchResponse, _a1 string, _a2 error) *StateInterface_ProcessBatchV2_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *StateInterface_ProcessBatchV2_Call) RunAndReturn(run func(context.Context, state.ProcessRequest, bool) (*state.ProcessBatchResponse, string, error)) *StateInterface_ProcessBatchV2_Call {
	_c.Call.Return(run)
	return _c
}

// ResetTrustedState provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateInterface) ResetTrustedState(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for ResetTrustedState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateInterface_ResetTrustedState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetTrustedState'
type StateInterface_ResetTrustedState_Call struct {
	*mock.Call
}

// ResetTrustedState is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) ResetTrustedState(ctx interface{}, batchNumber interface{}, dbTx interface{}) *StateInterface_ResetTrustedState_Call {
	return &StateInterface_ResetTrustedState_Call{Call: _e.mock.On("ResetTrustedState", ctx, batchNumber, dbTx)}
}

func (_c *StateInterface_ResetTrustedState_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx pgx.Tx)) *StateInterface_ResetTrustedState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_ResetTrustedState_Call) Return(_a0 error) *StateInterface_ResetTrustedState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateInterface_ResetTrustedState_Call) RunAndReturn(run func(context.Context, uint64, pgx.Tx) error) *StateInterface_ResetTrustedState_Call {
	_c.Call.Return(run)
	return _c
}

// StoreL2Block provides a mock function with given fields: ctx, batchNumber, l2Block, txsEGPLog, dbTx
func (_m *StateInterface) StoreL2Block(ctx context.Context, batchNumber uint64, l2Block *state.ProcessBlockResponse, txsEGPLog []*state.EffectiveGasPriceLog, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, batchNumber, l2Block, txsEGPLog, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for StoreL2Block")
	}

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *state.ProcessBlockResponse, []*state.EffectiveGasPriceLog, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, batchNumber, l2Block, txsEGPLog, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *state.ProcessBlockResponse, []*state.EffectiveGasPriceLog, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, batchNumber, l2Block, txsEGPLog, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, *state.ProcessBlockResponse, []*state.EffectiveGasPriceLog, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, l2Block, txsEGPLog, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInterface_StoreL2Block_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreL2Block'
type StateInterface_StoreL2Block_Call struct {
	*mock.Call
}

// StoreL2Block is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - l2Block *state.ProcessBlockResponse
//   - txsEGPLog []*state.EffectiveGasPriceLog
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) StoreL2Block(ctx interface{}, batchNumber interface{}, l2Block interface{}, txsEGPLog interface{}, dbTx interface{}) *StateInterface_StoreL2Block_Call {
	return &StateInterface_StoreL2Block_Call{Call: _e.mock.On("StoreL2Block", ctx, batchNumber, l2Block, txsEGPLog, dbTx)}
}

func (_c *StateInterface_StoreL2Block_Call) Run(run func(ctx context.Context, batchNumber uint64, l2Block *state.ProcessBlockResponse, txsEGPLog []*state.EffectiveGasPriceLog, dbTx pgx.Tx)) *StateInterface_StoreL2Block_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*state.ProcessBlockResponse), args[3].([]*state.EffectiveGasPriceLog), args[4].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_StoreL2Block_Call) Return(_a0 common.Hash, _a1 error) *StateInterface_StoreL2Block_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateInterface_StoreL2Block_Call) RunAndReturn(run func(context.Context, uint64, *state.ProcessBlockResponse, []*state.EffectiveGasPriceLog, pgx.Tx) (common.Hash, error)) *StateInterface_StoreL2Block_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWIPBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *StateInterface) UpdateWIPBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWIPBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateInterface_UpdateWIPBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWIPBatch'
type StateInterface_UpdateWIPBatch_Call struct {
	*mock.Call
}

// UpdateWIPBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - receipt state.ProcessingReceipt
//   - dbTx pgx.Tx
func (_e *StateInterface_Expecter) UpdateWIPBatch(ctx interface{}, receipt interface{}, dbTx interface{}) *StateInterface_UpdateWIPBatch_Call {
	return &StateInterface_UpdateWIPBatch_Call{Call: _e.mock.On("UpdateWIPBatch", ctx, receipt, dbTx)}
}

func (_c *StateInterface_UpdateWIPBatch_Call) Run(run func(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx)) *StateInterface_UpdateWIPBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(state.ProcessingReceipt), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterface_UpdateWIPBatch_Call) Return(_a0 error) *StateInterface_UpdateWIPBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateInterface_UpdateWIPBatch_Call) RunAndReturn(run func(context.Context, state.ProcessingReceipt, pgx.Tx) error) *StateInterface_UpdateWIPBatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewStateInterface creates a new instance of StateInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateInterface {
	mock := &StateInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}