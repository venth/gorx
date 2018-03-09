package operator

import "github.com/venth/gorx"

func (o *observable) OnErrorResumeWith(observable gorx.Observable) gorx.Observable {
	return o.OnErrorResumeNext(func(error) gorx.Observable {
		return observable
	})
}
