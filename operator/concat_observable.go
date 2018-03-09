package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/observer"
)

func Concat(observables ...gorx.Observable) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		Empty().Subscribe(
			newConcatenatingObserver(emissionObserver, 0, &observables),
		)
	})
}

func newConcatenatingObserver(emissionObserver gorx.Observer, current int, observables *[]gorx.Observable) gorx.Observer {
	return observer.NewDelegatingObserver(
		emissionObserver,
		func() {
			if current < len(*observables) {
				sequence := (*observables)[current]
				sequence.Subscribe(
					newConcatenatingObserver(emissionObserver, current+1, observables),
				)
			} else {
				emissionObserver.OnComplete()
			}
		},
	)
}
