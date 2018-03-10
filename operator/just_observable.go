package operator

import (
	"errors"
	"reflect"

	"github.com/venth/gorx"
)

func Just(elements ...interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		nilPassed := forEachElementNotifyJustObserver(&elements, subscriptionState, emissionObserver)

		if !(nilPassed || subscriptionState.IsDisposed()) {
			emissionObserver.OnComplete()
		}
	})
}

func forEachElementNotifyJustObserver(elements *[]interface{}, subscriptionState gorx.DisposableState, emissionObserver gorx.Observer) bool {
	completed := false
	nilPassed := false
	elementsLen := len(*elements)
	for idx := 0; idx < elementsLen && !completed; idx++ {
		element := (*elements)[idx]
		completed, nilPassed = notifyJustObserver(element, subscriptionState, emissionObserver)
	}
	return nilPassed
}

func notifyJustObserver(element interface{}, subscriptionState gorx.DisposableState, emissionObserver gorx.Observer) (bool, bool) {
	nilPassed := !reflect.ValueOf(element).IsValid()
	completed := subscriptionState.IsDisposed() || nilPassed
	if nilPassed {
		emissionObserver.OnError(errors.New("nil element passed to observable.Just"))
	} else if !completed {
		emissionObserver.OnNext(element)
	}
	return completed, nilPassed
}
