package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/observer"
)

func (o *observable) Map(mapFunc func(interface{}) interface{}) gorx.Observable {
	return &mappingObservable{
		mapFunc: mapFunc,
		Observable: o,
	}
}

type mappingObservable struct {
	mapFunc func(interface{}) interface{}
	gorx.Observable
}

func (o *mappingObservable) Subscribe(emissionObserver gorx.Observer) gorx.Disposable {
	var mappingOnNextFunc func(interface{})

	mappingOnNextFunc = newMappingOnNextFunc(emissionObserver, o.mapFunc)

	mappingObserver := observer.NewDelegatingObserver(emissionObserver, mappingOnNextFunc)
	return o.Observable.Subscribe(mappingObserver)
}

func newMappingOnNextFunc(emitter gorx.Observer, mapFunc func(interface{}) interface{}) func(interface{}) {
	return func(element interface{}) {
		emitter.OnNext(mapFunc(element))
	}
}