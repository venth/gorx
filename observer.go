package gorx


type UnboundObserver interface {
	Observer
	Bind(disposable Disposable) BoundObserver
}

type BoundObserver interface {
	Observer
	Unbind() UnboundObserver
}

type Observer interface {
	OnNext(element interface{})
	OnError(err error)
	OnComplete()
}
