package gorx

type ObservableSource interface{
	Subscribe(observer UnboundObserver) Subscription
}

type Observable interface {
	ObservableSource
}
