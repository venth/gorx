package operator

import (
	"reflect"

	"github.com/venth/gorx"
)

func FromSlice(elements interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		listVal := reflect.ValueOf(elements)
		completed, nilPassed := false, false
		for i := 0; i < listVal.Len() && !completed; i++ {
			element := listVal.Index(i).Interface()
			completed, nilPassed = notifyJustObserver(element, subscriptionState, emissionObserver)
		}

		if !(nilPassed || subscriptionState.IsDisposed()) {
			emissionObserver.OnComplete()
		}
	})
}
