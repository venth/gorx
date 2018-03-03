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
	subscribed, subscriptionRunner := newSubscribedObservable(o.emitSequence, emissionObserver)
	subscriptionRunner.Run()

	return subscribed
}
