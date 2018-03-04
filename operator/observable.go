package operator

import (
	"github.com/venth/gorx"
)

func newObservable(emitSequence gorx.EmitSequence) gorx.Observable {
	return &observable{
		emitSequence: emitSequence,
	}
}

type observable struct {
	emitSequence gorx.EmitSequence
}

func (o *observable) Subscribe(emissionObserver gorx.Observer) gorx.Disposable {
	subscription := newSubscribedObservable(o.emitSequence, emissionObserver)

	if subscribed, ok := subscription.(gorx.SubscribedObserver); ok {
		subscribed.OnSubscribe(o.emitSequence, subscription)
	}

	return subscription
}
