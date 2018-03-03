package gorx

type Disposable interface {
	Dispose()
	DisposableState
}

type DisposableState interface {
	IsDisposed() bool
}
