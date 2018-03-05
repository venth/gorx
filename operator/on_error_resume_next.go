package operator

import "github.com/venth/gorx"

func (o *observable) OnErrorResumeNext(resumeObservable gorx.Observable) gorx.Observable {
	return CreateObservable(func(emissionObserver gorx.Observer, subscriptionState gorx.DisposableState) {
		ob := &errorResumeNextObserver{
			resumeObservable: resumeObservable,
			Observer: emissionObserver,
		}

		o.Subscribe(ob)
	})
}

type errorResumeNextObserver struct {
	gorx.Observer
	resumeObservable gorx.Observable
}

func (o *errorResumeNextObserver) OnError(err error) {
	o.Observer.OnError(err)
}
