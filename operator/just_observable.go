package operator

import "github.com/venth/gorx"

func Just(elements ... interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		for element := range elements {
			if !subscriptionState.IsDisposed() {
				emissionObserver.OnNext(element)
			}
		}
		if !subscriptionState.IsDisposed() {
			emissionObserver.OnComplete()
		}
	})
}
