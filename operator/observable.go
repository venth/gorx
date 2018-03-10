package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/errors"
)

func newObservable(emitSequence gorx.EmitSequence) gorx.Observable {
	return &observable{
		emitSequence: emitSequence,
	}
}

type observable struct {
	emitSequence gorx.EmitSequence
}

func (o *observable) Subscribe(emissionObserver gorx.Observer) gorx.Disposable {
	subscription := newSubscribedObservable(o.emitSequence, emissionObserver)

	if subscribed, ok := subscription.(gorx.SubscribedObserver); ok {
		subscribed.OnSubscribe(o.emitSequence, subscription)
	}

	return subscription
}

func newSubscribedObservable(emitSequence gorx.EmitSequence, emissionObserver gorx.Observer) gorx.Disposable {
	subscribed := &subscribedObservable{
		emitSequence: emitSequence,
		Observer:     emissionObserver,
		complete:     make(chan struct{}),
	}

	return subscribed
}

type subscribedObservable struct {
	emitSequence gorx.EmitSequence
	gorx.Observer
	complete chan struct{}
}

func (s *subscribedObservable) OnNext(element interface{}) {
	if !s.IsDisposed() {
		s.Observer.OnNext(element)
	}
}

func (s *subscribedObservable) OnError(err error) {
	if !s.IsDisposed() {
		s.Observer.OnError(err)
		s.Dispose()
	}
}

func (s *subscribedObservable) OnComplete() {
	if !s.IsDisposed() {
		s.Observer.OnComplete()
		s.Dispose()
	}
}

func (s *subscribedObservable) Dispose() {
	errors.DontPanicCalling(func() {
		close(s.complete)
	})
}

func (s *subscribedObservable) IsDisposed() bool {
	select {
	case <-s.complete:
		return true
	default:
		return false
	}
}

func (s *subscribedObservable) OnSubscribe(emitSequence gorx.EmitSequence, subscription gorx.Disposable) {
	if subscribedObserver, ok := s.Observer.(gorx.SubscribedObserver); ok {
		subscribedObserver.OnSubscribe(s.emitSequence, subscription)
	}
	s.emitSequence(s.Observer, subscription)
}
