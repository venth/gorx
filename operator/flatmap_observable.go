package operator

import (
	"errors"
	"fmt"

	"github.com/venth/gorx"
	observableErrors "github.com/venth/gorx/errors"
)

func (o *observable) FlatMap(mapFunc func(interface{}) gorx.Observable) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		go o.Subscribe(newFlatMappingObserver(emissionObserver, mapFunc))
	})
}

func newFlatMappingObserver(emissionObserver gorx.Observer, mapFunc func(interface{}) gorx.Observable) gorx.Observer {
	return &flatMappingObserver{
		emissionObserver: emissionObserver,
		flatMappingFunc:  mapFunc,
		flattenSequence:  make(chan gorx.Observable),
	}
}

func (o *flatMappingObserver) OnNext(element interface{}) {
	flatMappedObservable := o.applyFlattingFunc(element)
	observableErrors.DontPanicCalling(func() {
		o.flattenSequence <- flatMappedObservable
	})
}
func (o *flatMappingObserver) applyFlattingFunc(element interface{}) gorx.Observable {
	var flatMappedObservable gorx.Observable
	defer func() {
		if r := recover(); r != nil {
			flatMappedObservable = Error(errors.New(fmt.Sprintf("An error occurred during flatMap function. %v", r)))
		}
	}()
	flatMappedObservable = o.flatMappingFunc(element)
	return flatMappedObservable
}

func (o *flatMappingObserver) OnError(err error) {
	errorObservable := Error(err)

	observableErrors.DontPanicCalling(func() {
		o.flattenSequence <- errorObservable
	})
}

func (o *flatMappingObserver) OnComplete() {
	close(o.flattenSequence)
}

func (o *flatMappingObserver) Run() {

}

type flatMappingObserver struct {
	emissionObserver gorx.Observer
	flatMappingFunc  func(interface{}) gorx.Observable
	complete         chan struct{}
	next             chan interface{}
	err              chan error
	flattenSequence  chan gorx.Observable
}
