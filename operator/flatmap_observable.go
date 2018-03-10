package operator

import (
	"fmt"

	"github.com/venth/gorx"
	"github.com/venth/gorx/observer"
)

func (o *observable) FlatMap(mapFunc func(interface{}) gorx.Observable) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		o.Subscribe(newFlatMappingObserver(emissionObserver, mapFunc))
	})
}

type noopDisposable struct{}

func (d *noopDisposable) Dispose() {}

func (d *noopDisposable) IsDisposed() bool {
	return false
}

func newFlatMappingObserver(emissionObserver gorx.Observer, mapFunc func(interface{}) gorx.Observable) gorx.Observer {
	flatMappingObserver := &flatMappingObserver{
		emissionObserver: emissionObserver,
		flatMappingFunc:  mapFunc,
		subscription:     &noopDisposable{},
	}
	mergingObserver := observer.NewObserver(
		func(el interface{}) {
			emissionObserver.OnNext(el)
		},
		func(err error) {
			emissionObserver.OnError(err)
			flatMappingObserver.Dispose()
		},
		func() {},
	)

	flatMappingObserver.mergingObserver = mergingObserver
	return flatMappingObserver
}

func (o *flatMappingObserver) Dispose() {
	o.subscription.Dispose()
}

func (o *flatMappingObserver) applyFlattingFunc(element interface{}) gorx.Observable {
	var flatMappedObservable gorx.Observable
	defer func() {
		if r := recover(); r != nil {
			flatMappedObservable = Error(fmt.Errorf("an error occurred during flatMap function; %v", r))
		}
	}()
	flatMappedObservable = o.flatMappingFunc(element)
	return flatMappedObservable
}

func (o *flatMappingObserver) OnNext(element interface{}) {
	flattenedObservable := o.applyFlattingFunc(element)
	flattenedObservable.Subscribe(o.mergingObserver)
}

func (o *flatMappingObserver) OnError(err error) {
	o.emissionObserver.OnError(err)
}

func (o *flatMappingObserver) OnComplete() {
	o.emissionObserver.OnComplete()
}

func (o *flatMappingObserver) OnSubscribe(emitSequence gorx.EmitSequence, subscription gorx.Disposable) {
	o.subscription = subscription
}

type flatMappingObserver struct {
	emissionObserver gorx.Observer
	mergingObserver  gorx.Observer
	flatMappingFunc  func(interface{}) gorx.Observable
	subscription     gorx.Disposable
}
