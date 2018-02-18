package disposable


type Disposable interface {
	Dispose() Disposed
	IsDisposed() bool
}

type Disposed interface {
	IsDisposed() bool
}
