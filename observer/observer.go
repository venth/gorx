package observer

import (
	"github.com/venth/gorx"
)

func NewDelegatingObserver(targetObserver gorx.Observer, handlers ...interface{}) gorx.Observer {
	return newObserver(
		delegatingOnNext(targetObserver),
		delegatingOnError(targetObserver),
		delegatingOnComplete(targetObserver),
		&handlers,
	)
}

func NewObserver(handlers ...interface{}) gorx.Observer {
	return newObserver(noopOnNext, noopOnError, noopOnComplete, &handlers)
}

func newObserver(nextFunc func(interface{}), errFunc func(error), completeFunc func(), handlers *[]interface{}) *observer {
	o := &observer{
		nextFunc:     nextFunc,
		errFunc:      errFunc,
		completeFunc: completeFunc,
	}

	for _, handler := range *handlers {
		switch handler := handler.(type) {
		case func():
			o.completeFunc = handler
			break

		case func(error):
			o.errFunc = handler
			break

		case func(interface{}):
			o.nextFunc = handler
			break
		}
	}
	return o
}

type observer struct {
	nextFunc     func(element interface{})
	errFunc      func(err error)
	completeFunc func()
}

func (o *observer) OnNext(element interface{}) {
	o.nextFunc(element)
}

func (o *observer) OnError(err error) {
	o.errFunc(err)
}

func (o *observer) OnComplete() {
	o.completeFunc()
}

func delegatingOnNext(observer gorx.Observer) func(interface{}) {
	return func(element interface{}) {
		observer.OnNext(element)
	}
}

func delegatingOnError(observer gorx.Observer) func(error) {
	return func(err error) {
		observer.OnError(err)
	}
}

func delegatingOnComplete(observer gorx.Observer) func() {
	return func() {
		observer.OnComplete()
	}
}

func noopOnNext(interface{}) {}
func noopOnError(error)      {}
func noopOnComplete()        {}
