package observer

import (
	"errors"
	"fmt"

	"github.com/venth/gorx"
	gorx_errors "github.com/venth/gorx/errors"
)

func NewQueuingObserver(emissionObserver gorx.Observer) gorx.Observer {
	return &queuingObserver{
		emissionObserver: emissionObserver,
		complete:         make(chan struct{}),
		next:             make(chan interface{}),
		err:              make(chan error),
	}
}

type queuingObserver struct {
	emissionObserver gorx.Observer
	complete         chan struct{}
	next             chan interface{}
	err              chan error
}

func (o *queuingObserver) Dispose() {
	gorx_errors.DontPanicCalling(func() {
		close(o.complete)
	})
	gorx_errors.DontPanicCalling(func() {
		close(o.err)
	})
	gorx_errors.DontPanicCalling(func() {
		close(o.next)
	})
}

func (o *queuingObserver) IsDisposed() bool {
	select {
	case <-o.complete:
		return true
	default:
		return false
	}
}

func (o *queuingObserver) OnNext(element interface{}) {
	o.next <- element
}

func (o *queuingObserver) OnError(err error) {
	o.err <- err
}

func (o *queuingObserver) OnComplete() {
	close(o.complete)
}

func (o *queuingObserver) Run() {
	defer o.Dispose()
	more := true
	emissionObserver := o.emissionObserver

	for more {
		select {
		case <-o.complete:
			emissionObserver.OnComplete()
			more = false
			break
		case err := <-o.err:
			more = o.notifyOnError(err)
			break
		case element := <-o.next:
			more = o.notifyOnNext(element)
		default:
		}
	}

}
func (o *queuingObserver) notifyOnNext(element interface{}) bool {
	more := true
	defer func() {
		if r := recover(); r != nil {
			more = o.notifyOnError(errors.New(fmt.Sprintf(
				"Recovered panic in function notifyOnNext for element: %v. %v", element, r)))
		}
	}()

	o.emissionObserver.OnNext(element)
	return more
}

func (o *queuingObserver) notifyOnError(err error) bool {
	o.emissionObserver.OnError(err)
	return false
}
