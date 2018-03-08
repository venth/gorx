package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/observer"
)

func Concat(observables ...gorx.Observable) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		completeSuppressionObserver := observer.NewDelegatingObserver(
			emissionObserver,
			func() {},
		)

		for idx := 0; idx < len(observables); idx++ {
			if subscriptionState.IsDisposed() {
				break
			}

			sequence := observables[idx]
			sequence.Subscribe(completeSuppressionObserver)
		}

		if !subscriptionState.IsDisposed() {
			emissionObserver.OnComplete()
		}
	})
}
