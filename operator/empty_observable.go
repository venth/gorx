package operator

import "github.com/venth/gorx"

func Empty() gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, state gorx.DisposableState) {
		emissionObserver.OnComplete()
	})
}
