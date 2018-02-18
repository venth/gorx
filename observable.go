package gorx

type ObservableSource interface{
	Subscribe(observer Observer) Subscription
}

type Observable interface {
	ObservableSource
}

type Subscription interface {
	Disposable
}
