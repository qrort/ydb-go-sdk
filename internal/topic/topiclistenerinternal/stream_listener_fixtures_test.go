package topiclistenerinternal

import (
	"sync/atomic"

	"github.com/rekby/fixenv"
	"github.com/rekby/fixenv/sf"
	"go.uber.org/mock/gomock"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/grpcwrapper/rawtopic/rawtopicreadermock"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/topic/topicreadercommon"
)

func StreamListener(e fixenv.Env) *streamListener {
	return listenerAndHandler(e).listener
}

func EventHandlerMock(e fixenv.Env) *MockEventHandler {
	return listenerAndHandler(e).handlerMock
}

func listenerAndHandler(e fixenv.Env) listenerAndHandlerPair {
	f := func() (*fixenv.GenericResult[listenerAndHandlerPair], error) {
		handler := NewMockEventHandler(MockController(e))

		listener := &streamListener{}
		listener.initVars(&atomic.Int64{})
		listener.stream = StreamMock(e)
		listener.streamClose = func(cause error) {}
		listener.handler = handler
		listener.sessions = PartitionStorage(e)

		return fixenv.NewGenericResult(listenerAndHandlerPair{
			handlerMock: handler,
			listener:    listener,
		}), nil
	}

	return fixenv.CacheResult(e, f)
}

type listenerAndHandlerPair struct {
	handlerMock *MockEventHandler
	listener    *streamListener
}

func PartitionStorage(e fixenv.Env) *topicreadercommon.PartitionSessionStorage {
	f := func() (*fixenv.GenericResult[*topicreadercommon.PartitionSessionStorage], error) {
		storage := &topicreadercommon.PartitionSessionStorage{}
		if err := storage.Add(PartitionSession(e)); err != nil {
			return nil, err
		}

		return fixenv.NewGenericResult(storage), nil
	}

	return fixenv.CacheResult(e, f)
}

func PartitionSession(e fixenv.Env) *topicreadercommon.PartitionSession {
	f := func() (*fixenv.GenericResult[*topicreadercommon.PartitionSession], error) {
		return fixenv.NewGenericResult(topicreadercommon.NewPartitionSession(
			sf.Context(e),
			"",
			0,
			0,
			"",
			2,
			0,
			0,
		)), nil
	}

	return fixenv.CacheResult(e, f)
}

func MockController(e fixenv.Env) *gomock.Controller {
	f := func() (*fixenv.GenericResult[*gomock.Controller], error) {
		mc := gomock.NewController(e.T().(gomock.TestReporter))

		return fixenv.NewGenericResult(mc), nil
	}

	return fixenv.CacheResult(e, f)
}

func StreamMock(e fixenv.Env) *rawtopicreadermock.MockTopicReaderStreamInterface {
	f := func() (*fixenv.GenericResult[*rawtopicreadermock.MockTopicReaderStreamInterface], error) {
		m := rawtopicreadermock.NewMockTopicReaderStreamInterface(MockController(e))

		return fixenv.NewGenericResult(m), nil
	}

	return fixenv.CacheResult(e, f)
}
