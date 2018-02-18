package gorx

import "github.com/venth/gorx/disposable"

type Observer interface {
	OnSubscribe(disposable disposable.Disposable)
	OnNext(element interface{})
	OnError()
	OnComplete()
}