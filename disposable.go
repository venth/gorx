package gorx


type Disposable interface {
	Dispose() Disposed
	IsDisposed() bool
}

type Disposed interface {
	IsDisposed() bool
}

