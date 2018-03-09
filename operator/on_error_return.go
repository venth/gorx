package operator

import "github.com/venth/gorx"

func (o *observable) OnErrorReturn(element interface{}) gorx.Observable {
	return o.OnErrorResumeNext(func(error) gorx.Observable {
		return Just(element)
	})
}
