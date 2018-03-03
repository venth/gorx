package operator

import (
	"github.com/venth/gorx"
)

func CreateObservable(emitSequence gorx.EmitSequence) gorx.Observable {
	return newObservable(emitSequence)
}
