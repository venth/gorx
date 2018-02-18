package gorx

type ObservableSource interface{
	Subscribe(observer Observer) Subscription
}

type Observable interface {
	ObservableSource
}
