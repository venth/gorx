package operator

import (
	"errors"
	"reflect"

	"github.com/venth/gorx"
)

func Just(elements ... interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		completedWithErrors := forEachElementNotifyObserver(emissionObserver, &elements, subscriptionState)

		if completedWithErrors && !subscriptionState.IsDisposed() {
			emissionObserver.OnComplete()
		}
	})
}

func forEachElementNotifyObserver(emissionObserver gorx.Observer, elements *[]interface{}, subscriptionState gorx.DisposableState) bool {
	completedWithErrors := true

	for idx := range *elements {
		element := (*elements)[idx]

		if subscriptionState.IsDisposed() {
			break
		}

		nilPassed := !reflect.ValueOf(element).IsValid()
		if nilPassed {
			emissionObserver.OnError(errors.New("nil element passed to observable.Just"))
			completedWithErrors = false
			break
		} else {
			emissionObserver.OnNext(element)
		}

	}
	return completedWithErrors
}
