package ydb

import (
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3/credentials"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/meta"
)

type customOptions struct {
	meta meta.Meta
}

type CustomOption func(opts *customOptions)

func WithCustomToken(accessToken string) CustomOption {
	return func(opts *customOptions) {
		opts.meta = opts.meta.WithCredentials(
			credentials.NewAccessTokenCredentials(
				accessToken,
				credentials.WithSourceInfo(fmt.Sprintf(`WithCustomToken("%s")`, accessToken)),
			),
		)
	}
}

func WithCustomCredentials(creds credentials.Credentials) CustomOption {
	return func(opts *customOptions) {
		opts.meta = opts.meta.WithCredentials(
			creds,
		)
	}
}

func WithCustomUserAgent(userAgent string) CustomOption {
	return func(opts *customOptions) {
		opts.meta = opts.meta.WithUserAgent(userAgent)
	}
}

func WithCustomDatabase(database string) CustomOption {
	return func(opts *customOptions) {
		opts.meta = opts.meta.WithDatabase(database)
	}
}
