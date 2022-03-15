package table

import (
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type Options struct {
	NoTraceErrors   []interface{}
	Idempotent      bool
	TxSettings      *TransactionSettings
	TxCommitOptions []options.CommitTransactionOption
	FastBackoff     retry.Backoff
	SlowBackoff     retry.Backoff
	Trace           trace.Table
}

type Option func(o *Options)

func WithIdempotent() Option {
	return func(o *Options) {
		o.Idempotent = true
	}
}

// WithNoTraceErrors provides management error wrapping with or without stacktrace points
func WithNoTraceErrors(noTraceErrors ...error) Option {
	return func(o *Options) {
		for i := range noTraceErrors {
			o.NoTraceErrors = append(o.NoTraceErrors, &noTraceErrors[i])
		}
	}
}

func WithTxSettings(tx *TransactionSettings) Option {
	return func(o *Options) {
		o.TxSettings = tx
	}
}

func WithTxCommitOptions(opts ...options.CommitTransactionOption) Option {
	return func(o *Options) {
		o.TxCommitOptions = append(o.TxCommitOptions, opts...)
	}
}

func WithTrace(t trace.Table) Option {
	return func(o *Options) {
		o.Trace = t
	}
}
