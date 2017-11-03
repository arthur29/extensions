// Automatically generated by MockGen. DO NOT EDIT!
// Source: mongo/interfaces/interfaces.go

package interfaces

import (
	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
)

// Mock of MongoDB interface
type MockMongoDB struct {
	ctrl     *gomock.Controller
	recorder *_MockMongoDBRecorder
}

// Recorder for MockMongoDB (not exported)
type _MockMongoDBRecorder struct {
	mock *MockMongoDB
}

func NewMockMongoDB(ctrl *gomock.Controller) *MockMongoDB {
	mock := &MockMongoDB{ctrl: ctrl}
	mock.recorder = &_MockMongoDBRecorder{mock}
	return mock
}

func (_m *MockMongoDB) EXPECT() *_MockMongoDBRecorder {
	return _m.recorder
}

func (_m *MockMongoDB) Run(cmd interface{}, result interface{}) error {
	ret := _m.ctrl.Call(_m, "Run", cmd, result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMongoDBRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}

func (_m *MockMongoDB) C(name string) (Collection, Session) {
	ret := _m.ctrl.Call(_m, "C", name)
	ret0, _ := ret[0].(Collection)
	ret1, _ := ret[1].(Session)
	return ret0, ret1
}

func (_mr *_MockMongoDBRecorder) C(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "C", arg0)
}

func (_m *MockMongoDB) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockMongoDBRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of Collection interface
type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *_MockCollectionRecorder
}

// Recorder for MockCollection (not exported)
type _MockCollectionRecorder struct {
	mock *MockCollection
}

func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	mock := &MockCollection{ctrl: ctrl}
	mock.recorder = &_MockCollectionRecorder{mock}
	return mock
}

func (_m *MockCollection) EXPECT() *_MockCollectionRecorder {
	return _m.recorder
}

func (_m *MockCollection) Find(query interface{}) Query {
	ret := _m.ctrl.Call(_m, "Find", query)
	ret0, _ := ret[0].(Query)
	return ret0
}

func (_mr *_MockCollectionRecorder) Find(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Find", arg0)
}

func (_m *MockCollection) Insert(docs ...interface{}) error {
	_s := []interface{}{}
	for _, _x := range docs {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Insert", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCollectionRecorder) Insert(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Insert", arg0...)
}

// Mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *_MockSessionRecorder
}

// Recorder for MockSession (not exported)
type _MockSessionRecorder struct {
	mock *MockSession
}

func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &_MockSessionRecorder{mock}
	return mock
}

func (_m *MockSession) EXPECT() *_MockSessionRecorder {
	return _m.recorder
}

func (_m *MockSession) Copy() *mgo_v2.Session {
	ret := _m.ctrl.Call(_m, "Copy")
	ret0, _ := ret[0].(*mgo_v2.Session)
	return ret0
}

func (_mr *_MockSessionRecorder) Copy() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Copy")
}

func (_m *MockSession) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockSessionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *_MockQueryRecorder
}

// Recorder for MockQuery (not exported)
type _MockQueryRecorder struct {
	mock *MockQuery
}

func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &_MockQueryRecorder{mock}
	return mock
}

func (_m *MockQuery) EXPECT() *_MockQueryRecorder {
	return _m.recorder
}

func (_m *MockQuery) Iter() Iter {
	ret := _m.ctrl.Call(_m, "Iter")
	ret0, _ := ret[0].(Iter)
	return ret0
}

func (_mr *_MockQueryRecorder) Iter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iter")
}

func (_m *MockQuery) All(result interface{}) error {
	ret := _m.ctrl.Call(_m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryRecorder) All(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "All", arg0)
}

// Mock of Iter interface
type MockIter struct {
	ctrl     *gomock.Controller
	recorder *_MockIterRecorder
}

// Recorder for MockIter (not exported)
type _MockIterRecorder struct {
	mock *MockIter
}

func NewMockIter(ctrl *gomock.Controller) *MockIter {
	mock := &MockIter{ctrl: ctrl}
	mock.recorder = &_MockIterRecorder{mock}
	return mock
}

func (_m *MockIter) EXPECT() *_MockIterRecorder {
	return _m.recorder
}

func (_m *MockIter) Next(result interface{}) bool {
	ret := _m.ctrl.Call(_m, "Next", result)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockIterRecorder) Next(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next", arg0)
}

func (_m *MockIter) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIterRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}