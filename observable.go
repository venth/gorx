package gorx

type ObservableSource interface {
	ConcatWith(observable Observable) Observable

	OnErrorResumeNext(resumeFunc func(err error) Observable) Observable
	OnErrorResumeWith(observable Observable) Observable
	OnErrorReturn(element interface{}) Observable

	FlatMap(mapFunc func(interface{}) Observable) Observable
	Map(mapFunc func(interface{}) interface{}) Observable

	Subscribe(emissionObserver Observer) Disposable
}

type Observable interface {
	ObservableSource
}
