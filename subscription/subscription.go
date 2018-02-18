package subscription

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/disposable"
)

type subscription struct {}

func New() gorx.Subscription {
	return &subscription{}
}

func (s *subscription) Dispose() gorx.Disposed {
	return disposable.NewDisposed()
}

func (s *subscription) IsDisposed() bool {
	return false
}
