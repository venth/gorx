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
	completed := false
	errorOccurred := false

	elementsLen := len(*elements)
	for idx := 0;  idx < elementsLen && !completed; idx++ {
		element := (*elements)[idx]

		errorOccurred = !reflect.ValueOf(element).IsValid()
		completed = subscriptionState.IsDisposed() || errorOccurred

		if errorOccurred {
			emissionObserver.OnError(errors.New("nil element passed to observable.Just"))
			errorOccurred = true
			completed = true
		} else {
			emissionObserver.OnNext(element)
		}
	}

	return errorOccurred
}
