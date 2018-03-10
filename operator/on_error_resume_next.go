package operator

import "github.com/venth/gorx"

func (o *observable) OnErrorResumeNext(resumeFunc func(err error) gorx.Observable) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		ob := &errorResumeNextObserver{
			resumeFunc:        resumeFunc,
			Observer:          emissionObserver,
			subscriptionState: subscriptionState,
		}

		o.Subscribe(ob)
	})
}

type errorResumeNextObserver struct {
	gorx.Observer
	resumeFunc        func(err error) gorx.Observable
	subscriptionState gorx.DisposableState
}

func (o *errorResumeNextObserver) OnError(err error) {
	resumedObservable := o.resumeFunc(err)
	resumedObservable.Subscribe(o.Observer)
}
