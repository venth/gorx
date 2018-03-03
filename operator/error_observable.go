package operator

import (
	"github.com/venth/gorx"
)

func Error(err error) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		if !subscriptionState.IsDisposed() {
			emissionObserver.OnError(err)
		}
	})
}
