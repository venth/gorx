package observable

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/subscription"
)

type observable struct {
	emitter gorx.Emitter
}

func New(emitter gorx.Emitter) gorx.Observable {
	return &observable{emitter: emitter}
}

func (o *observable) Subscribe(observer gorx.UnboundObserver) gorx.Subscription {
	return subscription.NewQueuingSubscription(o.emitter, observer)
}
