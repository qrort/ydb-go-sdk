// Code generated by MockGen. DO NOT EDIT.
// Source: batched_stream_reader_interface.go
//
// Generated by this command:
//
//	mockgen -source batched_stream_reader_interface.go --typed -destination batched_stream_reader_mock_test.go -package topicreaderinternal -write_package_comment=false
package topicreaderinternal

import (
	context "context"
	reflect "reflect"

	topicreadercommon "github.com/ydb-platform/ydb-go-sdk/v3/internal/topic/topicreadercommon"
	tx "github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	gomock "go.uber.org/mock/gomock"
)

// MockbatchedStreamReader is a mock of batchedStreamReader interface.
type MockbatchedStreamReader struct {
	ctrl     *gomock.Controller
	recorder *MockbatchedStreamReaderMockRecorder
}

// MockbatchedStreamReaderMockRecorder is the mock recorder for MockbatchedStreamReader.
type MockbatchedStreamReaderMockRecorder struct {
	mock *MockbatchedStreamReader
}

// NewMockbatchedStreamReader creates a new mock instance.
func NewMockbatchedStreamReader(ctrl *gomock.Controller) *MockbatchedStreamReader {
	mock := &MockbatchedStreamReader{ctrl: ctrl}
	mock.recorder = &MockbatchedStreamReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbatchedStreamReader) EXPECT() *MockbatchedStreamReaderMockRecorder {
	return m.recorder
}

// CloseWithError mocks base method.
func (m *MockbatchedStreamReader) CloseWithError(ctx context.Context, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", ctx, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockbatchedStreamReaderMockRecorder) CloseWithError(ctx, err any) *MockbatchedStreamReaderCloseWithErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockbatchedStreamReader)(nil).CloseWithError), ctx, err)
	return &MockbatchedStreamReaderCloseWithErrorCall{Call: call}
}

// MockbatchedStreamReaderCloseWithErrorCall wrap *gomock.Call
type MockbatchedStreamReaderCloseWithErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderCloseWithErrorCall) Return(arg0 error) *MockbatchedStreamReaderCloseWithErrorCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderCloseWithErrorCall) Do(f func(context.Context, error) error) *MockbatchedStreamReaderCloseWithErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderCloseWithErrorCall) DoAndReturn(f func(context.Context, error) error) *MockbatchedStreamReaderCloseWithErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Commit mocks base method.
func (m *MockbatchedStreamReader) Commit(ctx context.Context, commitRange topicreadercommon.CommitRange) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, commitRange)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockbatchedStreamReaderMockRecorder) Commit(ctx, commitRange any) *MockbatchedStreamReaderCommitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockbatchedStreamReader)(nil).Commit), ctx, commitRange)
	return &MockbatchedStreamReaderCommitCall{Call: call}
}

// MockbatchedStreamReaderCommitCall wrap *gomock.Call
type MockbatchedStreamReaderCommitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderCommitCall) Return(arg0 error) *MockbatchedStreamReaderCommitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderCommitCall) Do(f func(context.Context, topicreadercommon.CommitRange) error) *MockbatchedStreamReaderCommitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderCommitCall) DoAndReturn(f func(context.Context, topicreadercommon.CommitRange) error) *MockbatchedStreamReaderCommitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PopBatchTx mocks base method.
func (m *MockbatchedStreamReader) PopBatchTx(ctx context.Context, tx tx.Notificator, opts ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopBatchTx", ctx, tx, opts)
	ret0, _ := ret[0].(*topicreadercommon.PublicBatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PopBatchTx indicates an expected call of PopBatchTx.
func (mr *MockbatchedStreamReaderMockRecorder) PopBatchTx(ctx, tx, opts any) *MockbatchedStreamReaderPopBatchTxCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopBatchTx", reflect.TypeOf((*MockbatchedStreamReader)(nil).PopBatchTx), ctx, tx, opts)
	return &MockbatchedStreamReaderPopBatchTxCall{Call: call}
}

// MockbatchedStreamReaderPopBatchTxCall wrap *gomock.Call
type MockbatchedStreamReaderPopBatchTxCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderPopBatchTxCall) Return(arg0 *topicreadercommon.PublicBatch, arg1 error) *MockbatchedStreamReaderPopBatchTxCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderPopBatchTxCall) Do(f func(context.Context, tx.Notificator, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderPopBatchTxCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderPopBatchTxCall) DoAndReturn(f func(context.Context, tx.Notificator, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderPopBatchTxCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadMessageBatch mocks base method.
func (m *MockbatchedStreamReader) ReadMessageBatch(ctx context.Context, opts ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessageBatch", ctx, opts)
	ret0, _ := ret[0].(*topicreadercommon.PublicBatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessageBatch indicates an expected call of ReadMessageBatch.
func (mr *MockbatchedStreamReaderMockRecorder) ReadMessageBatch(ctx, opts any) *MockbatchedStreamReaderReadMessageBatchCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessageBatch", reflect.TypeOf((*MockbatchedStreamReader)(nil).ReadMessageBatch), ctx, opts)
	return &MockbatchedStreamReaderReadMessageBatchCall{Call: call}
}

// MockbatchedStreamReaderReadMessageBatchCall wrap *gomock.Call
type MockbatchedStreamReaderReadMessageBatchCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderReadMessageBatchCall) Return(arg0 *topicreadercommon.PublicBatch, arg1 error) *MockbatchedStreamReaderReadMessageBatchCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderReadMessageBatchCall) Do(f func(context.Context, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderReadMessageBatchCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderReadMessageBatchCall) DoAndReturn(f func(context.Context, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderReadMessageBatchCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WaitInit mocks base method.
func (m *MockbatchedStreamReader) WaitInit(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitInit", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitInit indicates an expected call of WaitInit.
func (mr *MockbatchedStreamReaderMockRecorder) WaitInit(ctx any) *MockbatchedStreamReaderWaitInitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitInit", reflect.TypeOf((*MockbatchedStreamReader)(nil).WaitInit), ctx)
	return &MockbatchedStreamReaderWaitInitCall{Call: call}
}

// MockbatchedStreamReaderWaitInitCall wrap *gomock.Call
type MockbatchedStreamReaderWaitInitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderWaitInitCall) Return(arg0 error) *MockbatchedStreamReaderWaitInitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderWaitInitCall) Do(f func(context.Context) error) *MockbatchedStreamReaderWaitInitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderWaitInitCall) DoAndReturn(f func(context.Context) error) *MockbatchedStreamReaderWaitInitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
