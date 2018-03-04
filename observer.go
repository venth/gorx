package gorx

//go:generate mockgen -source=./observer.go -destination=./observer_mock.go -package gorx github.com/venth/gorx

type Observer interface {
	OnNext(element interface{})
	OnError(err error)
	OnComplete()
}

type SubscribedObserver interface {
	OnSubscribe(emitSequence EmitSequence, subscription Disposable)
}

type EmitSequence func(Observer, DisposableState)
