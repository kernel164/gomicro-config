package ssm

import (
	"errors"

	"github.com/micro/go-micro/config/source"
)

type ssmWatcher struct {
	s    *ssmSource
	exit chan bool
}

func newWatcher(s *ssmSource) (*ssmWatcher, error) {
	return &ssmWatcher{
		s:    s,
		exit: make(chan bool),
	}, nil
}

func (w *ssmWatcher) Next() (*source.ChangeSet, error) {
	<-w.exit
	return nil, errors.New("ssm watcher stopped")
}

func (w *ssmWatcher) Stop() error {
	select {
	case <-w.exit:
	default:
		// TODO: stop water resource.
	}
	return nil
}
