// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/celestiaorg/celestia-node/nodebuilder/state (interfaces: Module)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
	types1 "github.com/tendermint/tendermint/types"

	namespace "github.com/celestiaorg/nmt/namespace"
)

// MockModule is a mock of Module interface.
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule.
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance.
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// AccountAddress mocks base method.
func (m *MockModule) AccountAddress(arg0 context.Context) (types.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountAddress", arg0)
	ret0, _ := ret[0].(types.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountAddress indicates an expected call of AccountAddress.
func (mr *MockModuleMockRecorder) AccountAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountAddress", reflect.TypeOf((*MockModule)(nil).AccountAddress), arg0)
}

// Balance mocks base method.
func (m *MockModule) Balance(arg0 context.Context) (*types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance", arg0)
	ret0, _ := ret[0].(*types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance.
func (mr *MockModuleMockRecorder) Balance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockModule)(nil).Balance), arg0)
}

// BalanceForAddress mocks base method.
func (m *MockModule) BalanceForAddress(arg0 context.Context, arg1 types.Address) (*types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceForAddress", arg0, arg1)
	ret0, _ := ret[0].(*types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceForAddress indicates an expected call of BalanceForAddress.
func (mr *MockModuleMockRecorder) BalanceForAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceForAddress", reflect.TypeOf((*MockModule)(nil).BalanceForAddress), arg0, arg1)
}

// BeginRedelegate mocks base method.
func (m *MockModule) BeginRedelegate(arg0 context.Context, arg1, arg2 types.ValAddress, arg3, arg4 math.Int, arg5 uint64) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginRedelegate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginRedelegate indicates an expected call of BeginRedelegate.
func (mr *MockModuleMockRecorder) BeginRedelegate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginRedelegate", reflect.TypeOf((*MockModule)(nil).BeginRedelegate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CancelUnbondingDelegation mocks base method.
func (m *MockModule) CancelUnbondingDelegation(arg0 context.Context, arg1 types.ValAddress, arg2, arg3, arg4 math.Int, arg5 uint64) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelUnbondingDelegation", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelUnbondingDelegation indicates an expected call of CancelUnbondingDelegation.
func (mr *MockModuleMockRecorder) CancelUnbondingDelegation(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelUnbondingDelegation", reflect.TypeOf((*MockModule)(nil).CancelUnbondingDelegation), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Delegate mocks base method.
func (m *MockModule) Delegate(arg0 context.Context, arg1 types.ValAddress, arg2, arg3 math.Int, arg4 uint64) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delegate indicates an expected call of Delegate.
func (mr *MockModuleMockRecorder) Delegate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegate", reflect.TypeOf((*MockModule)(nil).Delegate), arg0, arg1, arg2, arg3, arg4)
}

// IsStopped mocks base method.
func (m *MockModule) IsStopped(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStopped", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStopped indicates an expected call of IsStopped.
func (mr *MockModuleMockRecorder) IsStopped(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStopped", reflect.TypeOf((*MockModule)(nil).IsStopped), arg0)
}

// QueryDelegation mocks base method.
func (m *MockModule) QueryDelegation(arg0 context.Context, arg1 types.ValAddress) (*types0.QueryDelegationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryDelegation", arg0, arg1)
	ret0, _ := ret[0].(*types0.QueryDelegationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDelegation indicates an expected call of QueryDelegation.
func (mr *MockModuleMockRecorder) QueryDelegation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryDelegation", reflect.TypeOf((*MockModule)(nil).QueryDelegation), arg0, arg1)
}

// QueryRedelegations mocks base method.
func (m *MockModule) QueryRedelegations(arg0 context.Context, arg1, arg2 types.ValAddress) (*types0.QueryRedelegationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRedelegations", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types0.QueryRedelegationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRedelegations indicates an expected call of QueryRedelegations.
func (mr *MockModuleMockRecorder) QueryRedelegations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRedelegations", reflect.TypeOf((*MockModule)(nil).QueryRedelegations), arg0, arg1, arg2)
}

// QueryUnbonding mocks base method.
func (m *MockModule) QueryUnbonding(arg0 context.Context, arg1 types.ValAddress) (*types0.QueryUnbondingDelegationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUnbonding", arg0, arg1)
	ret0, _ := ret[0].(*types0.QueryUnbondingDelegationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUnbonding indicates an expected call of QueryUnbonding.
func (mr *MockModuleMockRecorder) QueryUnbonding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUnbonding", reflect.TypeOf((*MockModule)(nil).QueryUnbonding), arg0, arg1)
}

// SubmitPayForData mocks base method.
func (m *MockModule) SubmitPayForData(arg0 context.Context, arg1 namespace.ID, arg2 []byte, arg3 math.Int, arg4 uint64) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitPayForData", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitPayForData indicates an expected call of SubmitPayForData.
func (mr *MockModuleMockRecorder) SubmitPayForData(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitPayForData", reflect.TypeOf((*MockModule)(nil).SubmitPayForData), arg0, arg1, arg2, arg3, arg4)
}

// SubmitTx mocks base method.
func (m *MockModule) SubmitTx(arg0 context.Context, arg1 types1.Tx) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTx", arg0, arg1)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTx indicates an expected call of SubmitTx.
func (mr *MockModuleMockRecorder) SubmitTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTx", reflect.TypeOf((*MockModule)(nil).SubmitTx), arg0, arg1)
}

// Transfer mocks base method.
func (m *MockModule) Transfer(arg0 context.Context, arg1 types.AccAddress, arg2, arg3 math.Int, arg4 uint64) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockModuleMockRecorder) Transfer(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockModule)(nil).Transfer), arg0, arg1, arg2, arg3, arg4)
}

// Undelegate mocks base method.
func (m *MockModule) Undelegate(arg0 context.Context, arg1 types.ValAddress, arg2, arg3 math.Int, arg4 uint64) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Undelegate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Undelegate indicates an expected call of Undelegate.
func (mr *MockModuleMockRecorder) Undelegate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Undelegate", reflect.TypeOf((*MockModule)(nil).Undelegate), arg0, arg1, arg2, arg3, arg4)
}
