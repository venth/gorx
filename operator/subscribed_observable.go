package operator

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/errors"
	"github.com/venth/gorx/observer"
)

func newSubscribedObservable(emitSequence gorx.EmitSequence, emissionObserver gorx.Observer) (gorx.Disposable, gorx.Runnable) {
	subscribed := &subscribedObservable{
		emitSequence:     emitSequence,
		emissionObserver: emissionObserver,
		complete:         make(chan struct{}),
	}
	return subscribed, subscribed
}

type subscribedObservable struct {
	emitSequence     gorx.EmitSequence
	emissionObserver gorx.Observer
	complete         chan struct{}
}

func (s *subscribedObservable) Run() {
	closingObserver := s.newDisposingOnFinalStateObserver(s.emissionObserver)
	s.emitSequence(closingObserver, s)
}
func (s *subscribedObservable) newDisposingOnFinalStateObserver(emissionObserver gorx.Observer) gorx.Observer {
	return observer.NewDelegatingObserver(
		emissionObserver,
		func(err error) {
			emissionObserver.OnError(err)
			s.Dispose()
		},
		func() {
			emissionObserver.OnComplete()
			s.Dispose()
		},
	)
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
