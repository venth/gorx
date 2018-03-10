package operator

import (
	"reflect"

	"github.com/venth/gorx"
)

func FromSlice(elements interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		elementList := reflect.ValueOf(elements)
		nilPassed := forEachSliceElement(elementList, subscriptionState, emissionObserver)

		if !(nilPassed || subscriptionState.IsDisposed()) {
			emissionObserver.OnComplete()
		}
	})
}
func forEachSliceElement(listVal reflect.Value, subscriptionState gorx.DisposableState, emissionObserver gorx.Observer) bool {
	completed, nilPassed := false, false
	for i := 0; i < listVal.Len() && !completed; i++ {
		element := listVal.Index(i).Interface()
		completed, nilPassed = notifyJustObserver(element, subscriptionState, emissionObserver)
	}
	return nilPassed
}
