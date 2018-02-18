package observable

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/subscription"
)

type observable struct {}

func Empty() gorx.Observable {
	return &observable{}
}

func (o *observable) Subscribe(observer gorx.Observer) gorx.Subscription {
	return subscription.New()
}
