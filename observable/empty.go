package observable

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/disposable"
)

type emptyObservable struct {}
type subscribedObserver struct {
	observer gorx.Observer
}

func Empty() gorx.Observable {
	return &emptyObservable{}
}

func (o *emptyObservable) Subscribe(observer gorx.UnboundObserver) gorx.Subscription {
	subscribed := &subscribedObserver{observer: observer}
	subscribed.run()
	return subscribed
}

func (o *subscribedObserver) Dispose() gorx.Disposed {
	return disposable.NewDisposed()
}

func (o *subscribedObserver) run() {
	o.OnComplete()
}

func (o *subscribedObserver) IsDisposed() bool {
	return false
}

func (o *subscribedObserver) OnNext(element interface{}) {
	o.observer.OnNext(element)
}

func (o *subscribedObserver) OnComplete() {
	o.observer.OnComplete()
}

func (o *subscribedObserver) OnError(err error) {
	o.observer.OnError(err)
}
