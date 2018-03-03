package gorx

type ObservableSource interface{
	FlatMap(mapFunc func(interface{}) Observable) Observable
	Map(mapFunc func(interface{}) interface{}) Observable
	Subscribe(emissionObserver Observer) Disposable
}

type Observable interface {
	ObservableSource
}
