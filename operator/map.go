package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/observer"
)

func (o *observable) Map(mapFunc func(interface{}) interface{}) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		mappingOnNextFunc := newMappingOnNextFunc(emissionObserver, mapFunc)

		mappingObserver := observer.NewDelegatingObserver(emissionObserver, mappingOnNextFunc)
		o.Subscribe(mappingObserver)
	})
}

func newMappingOnNextFunc(emitter gorx.Observer, mapFunc func(interface{}) interface{}) func(interface{}) {
	return func(element interface{}) {
		emitter.OnNext(mapFunc(element))
	}
}
