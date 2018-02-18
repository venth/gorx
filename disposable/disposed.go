package disposable

import "github.com/venth/gorx"

type disposed struct {}

func NewDisposed() gorx.Disposed {
	return &disposed{}
}

func (d *disposed) IsDisposed() bool {
	return true
}