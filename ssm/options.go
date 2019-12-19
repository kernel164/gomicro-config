package ssm

import (
	"context"

	"github.com/micro/go-micro/config/source"
)

type ssmParameterKeyKey struct{}
type ssmParameterEncryptedKey struct{}

func WithSSMParameterKey(ssmParameterKey string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, ssmParameterKeyKey{}, ssmParameterKey)
	}
}

func WithSSMParameterEncrypted(ssmParameterEncrypted bool) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, ssmParameterEncryptedKey{}, ssmParameterEncrypted)
	}
}
