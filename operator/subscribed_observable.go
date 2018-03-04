package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/errors"
)

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
	complete     chan struct{}
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
	} else {
		s.emitSequence(s.Observer, subscription)
	}
}
