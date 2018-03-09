package operator

import "github.com/venth/gorx"

func (o *observable) ConcatWith(observable gorx.Observable) gorx.Observable {
	return Concat(o, observable)
}
