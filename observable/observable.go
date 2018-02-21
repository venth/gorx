package observable

import (
	"github.com/venth/gorx"
	"github.com/venth/gorx/subscription"
	"golang.org/x/net/html/atom"
)

type observable struct {
	emitter gorx.Emitter
}

func New(emitter gorx.Emitter) gorx.Observable {
	return &observable{emitter: emitter}
}

func (o *observable) Subscribe(observer gorx.Observer) gorx.Subscription {
	subscribed := subscription.NewQueuingSubscription(o.emitter, observer)
	subscribed.Run()

	//XXX correct implementation
	return subscribed
}
