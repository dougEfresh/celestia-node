// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/celestiaorg/celestia-node/nodebuilder/p2p (interfaces: Module)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	network "github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
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

// BandwidthForPeer mocks base method.
func (m *MockModule) BandwidthForPeer(arg0 context.Context, arg1 peer.ID) metrics.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BandwidthForPeer", arg0, arg1)
	ret0, _ := ret[0].(metrics.Stats)
	return ret0
}

// BandwidthForPeer indicates an expected call of BandwidthForPeer.
func (mr *MockModuleMockRecorder) BandwidthForPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BandwidthForPeer", reflect.TypeOf((*MockModule)(nil).BandwidthForPeer), arg0, arg1)
}

// BandwidthForProtocol mocks base method.
func (m *MockModule) BandwidthForProtocol(arg0 context.Context, arg1 protocol.ID) metrics.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BandwidthForProtocol", arg0, arg1)
	ret0, _ := ret[0].(metrics.Stats)
	return ret0
}

// BandwidthForProtocol indicates an expected call of BandwidthForProtocol.
func (mr *MockModuleMockRecorder) BandwidthForProtocol(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BandwidthForProtocol", reflect.TypeOf((*MockModule)(nil).BandwidthForProtocol), arg0, arg1)
}

// BandwidthStats mocks base method.
func (m *MockModule) BandwidthStats(arg0 context.Context) metrics.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BandwidthStats", arg0)
	ret0, _ := ret[0].(metrics.Stats)
	return ret0
}

// BandwidthStats indicates an expected call of BandwidthStats.
func (mr *MockModuleMockRecorder) BandwidthStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BandwidthStats", reflect.TypeOf((*MockModule)(nil).BandwidthStats), arg0)
}

// BlockPeer mocks base method.
func (m *MockModule) BlockPeer(arg0 context.Context, arg1 peer.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockPeer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlockPeer indicates an expected call of BlockPeer.
func (mr *MockModuleMockRecorder) BlockPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockPeer", reflect.TypeOf((*MockModule)(nil).BlockPeer), arg0, arg1)
}

// ClosePeer mocks base method.
func (m *MockModule) ClosePeer(arg0 context.Context, arg1 peer.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosePeer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClosePeer indicates an expected call of ClosePeer.
func (mr *MockModuleMockRecorder) ClosePeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePeer", reflect.TypeOf((*MockModule)(nil).ClosePeer), arg0, arg1)
}

// Connect mocks base method.
func (m *MockModule) Connect(arg0 context.Context, arg1 peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockModuleMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockModule)(nil).Connect), arg0, arg1)
}

// Connectedness mocks base method.
func (m *MockModule) Connectedness(arg0 context.Context, arg1 peer.ID) network.Connectedness {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connectedness", arg0, arg1)
	ret0, _ := ret[0].(network.Connectedness)
	return ret0
}

// Connectedness indicates an expected call of Connectedness.
func (mr *MockModuleMockRecorder) Connectedness(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connectedness", reflect.TypeOf((*MockModule)(nil).Connectedness), arg0, arg1)
}

// Info mocks base method.
func (m *MockModule) Info(arg0 context.Context) peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(peer.AddrInfo)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockModuleMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockModule)(nil).Info), arg0)
}

// IsProtected mocks base method.
func (m *MockModule) IsProtected(arg0 context.Context, arg1 peer.ID, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProtected", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsProtected indicates an expected call of IsProtected.
func (mr *MockModuleMockRecorder) IsProtected(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProtected", reflect.TypeOf((*MockModule)(nil).IsProtected), arg0, arg1, arg2)
}

// ListBlockedPeers mocks base method.
func (m *MockModule) ListBlockedPeers(arg0 context.Context) []peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlockedPeers", arg0)
	ret0, _ := ret[0].([]peer.ID)
	return ret0
}

// ListBlockedPeers indicates an expected call of ListBlockedPeers.
func (mr *MockModuleMockRecorder) ListBlockedPeers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockedPeers", reflect.TypeOf((*MockModule)(nil).ListBlockedPeers), arg0)
}

// NATStatus mocks base method.
func (m *MockModule) NATStatus(arg0 context.Context) (network.Reachability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NATStatus", arg0)
	ret0, _ := ret[0].(network.Reachability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NATStatus indicates an expected call of NATStatus.
func (mr *MockModuleMockRecorder) NATStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NATStatus", reflect.TypeOf((*MockModule)(nil).NATStatus), arg0)
}

// PeerInfo mocks base method.
func (m *MockModule) PeerInfo(arg0 context.Context, arg1 peer.ID) peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerInfo", arg0, arg1)
	ret0, _ := ret[0].(peer.AddrInfo)
	return ret0
}

// PeerInfo indicates an expected call of PeerInfo.
func (mr *MockModuleMockRecorder) PeerInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerInfo", reflect.TypeOf((*MockModule)(nil).PeerInfo), arg0, arg1)
}

// Peers mocks base method.
func (m *MockModule) Peers(arg0 context.Context) []peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers", arg0)
	ret0, _ := ret[0].([]peer.ID)
	return ret0
}

// Peers indicates an expected call of Peers.
func (mr *MockModuleMockRecorder) Peers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockModule)(nil).Peers), arg0)
}

// Protect mocks base method.
func (m *MockModule) Protect(arg0 context.Context, arg1 peer.ID, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Protect", arg0, arg1, arg2)
}

// Protect indicates an expected call of Protect.
func (mr *MockModuleMockRecorder) Protect(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protect", reflect.TypeOf((*MockModule)(nil).Protect), arg0, arg1, arg2)
}

// PubSubPeers mocks base method.
func (m *MockModule) PubSubPeers(arg0 context.Context, arg1 string) []peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubPeers", arg0, arg1)
	ret0, _ := ret[0].([]peer.ID)
	return ret0
}

// PubSubPeers indicates an expected call of PubSubPeers.
func (mr *MockModuleMockRecorder) PubSubPeers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubPeers", reflect.TypeOf((*MockModule)(nil).PubSubPeers), arg0, arg1)
}

// UnblockPeer mocks base method.
func (m *MockModule) UnblockPeer(arg0 context.Context, arg1 peer.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnblockPeer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnblockPeer indicates an expected call of UnblockPeer.
func (mr *MockModuleMockRecorder) UnblockPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnblockPeer", reflect.TypeOf((*MockModule)(nil).UnblockPeer), arg0, arg1)
}

// Unprotect mocks base method.
func (m *MockModule) Unprotect(arg0 context.Context, arg1 peer.ID, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unprotect", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Unprotect indicates an expected call of Unprotect.
func (mr *MockModuleMockRecorder) Unprotect(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unprotect", reflect.TypeOf((*MockModule)(nil).Unprotect), arg0, arg1, arg2)
}
