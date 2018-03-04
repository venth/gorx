package operator

import (
	"errors"
	"reflect"

	"github.com/venth/gorx"
)

func Just(elements ... interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		completedWithErrors := true

		for idx := 0; idx < len(elements); idx++ {
			element := elements[idx]
			if subscriptionState.IsDisposed() {
				break
			}

			if reflect.ValueOf(element).IsValid() {
				emissionObserver.OnNext(element)
			} else {
				emissionObserver.OnError(errors.New("nil element passed to observable.Just"))
				completedWithErrors = false
				break
			}

		}
		if completedWithErrors && !subscriptionState.IsDisposed() {
			emissionObserver.OnComplete()
		}
	})
}
