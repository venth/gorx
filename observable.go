package gorx

type ObservableSource interface{
	OnErrorResumeNext(resumeFunc func(err error) Observable) Observable
	FlatMap(mapFunc func(interface{}) Observable) Observable
	Map(mapFunc func(interface{}) interface{}) Observable
	Subscribe(emissionObserver Observer) Disposable
}

type Observable interface {
	ObservableSource
}
