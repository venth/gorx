package subscription

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/disposable"
	"github.com/venth/gorx/observer"
)


type queuingSubscription struct {
	queuingObserver gorx.UnboundObserver
}

func NewQueuingSubscription(emitter gorx.Emitter, observerToDecorate gorx.UnboundObserver) gorx.Subscription {

	queuingObserver := observer.NewQueuingObserver(emitter, observerToDecorate)
	subscription := &queuingSubscription{queuingObserver: queuingObserver}
	queuingObserver.Bind(subscription)

	return subscription
}

func (s *queuingSubscription) Dispose() gorx.Disposed {
	return disposable.NewDisposed()
}

func (s *queuingSubscription) IsDisposed() bool {
	return false
}
