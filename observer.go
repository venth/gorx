package gorx

type Observer interface {
	OnSubscribe(disposable Disposable)
	OnNext(element interface{})
	OnError()
	OnComplete()
}