package gorx


//go:generate mockgen -source=./observer.go -destination=./observer_mock.go -package gorx github.com/venth/gorx

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
