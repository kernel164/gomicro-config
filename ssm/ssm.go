// Package ssm loads changesets from a ssm
package ssm

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/micro/go-micro/config/source"
)

type ssmSource struct {
	ssm                   *ssm.SSM
	ssmParameterKey       string
	ssmParameterEncrypted bool
	opts                  source.Options
	cs                    *source.ChangeSet
}

func (s *ssmSource) read() (*source.ChangeSet, error) {
	param, err := s.ssm.GetParameter(&ssm.GetParameterInput{
		Name:           &s.ssmParameterKey,
		WithDecryption: &s.ssmParameterEncrypted,
	})
	if err != nil {
		return nil, err
	}
	b := []byte(*param.Parameter.Value)
	ft := s.opts.Encoder.String()
	cs := &source.ChangeSet{
		Data:      b,
		Format:    ft,
		Timestamp: time.Now(),
		Source:    s.String(),
	}
	cs.Checksum = cs.Sum()
	return cs, err
}

func (s *ssmSource) Read() (*source.ChangeSet, error) {
	cs, err := s.read()
	if err != nil {
		return nil, err
	}
	s.cs = cs
	return cs, nil
}

func (s *ssmSource) Watch() (source.Watcher, error) {
	return newWatcher(s)
}

func (s *ssmSource) String() string {
	return "ssm"
}

func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	ssmParameterKey, ok := options.Context.Value(ssmParameterKeyKey{}).(string)
	if !ok || ssmParameterKey == "" {
		panic("ssm parameter key option invalid or missing.")
	}
	ssmParameterEncrypted, _ := options.Context.Value(ssmParameterEncryptedKey{}).(bool)
	ssm := ssm.New(session.New())
	return &ssmSource{
		ssm:                   ssm,
		ssmParameterKey:       ssmParameterKey,
		ssmParameterEncrypted: ssmParameterEncrypted,
		opts:                  options,
	}
}
