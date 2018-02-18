package gorx

import "github.com/venth/gorx/disposable"

type ObservableSource interface{
	Subscribe(observer Observer) ObservedSource
}

type ObservedSource interface {
	disposable.Disposable
}
