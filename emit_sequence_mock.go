package gorx

type notifiable interface {
	notify(observer Observer)
}

type nextEmission struct {
	element interface{}
}

func (e *nextEmission) notify(observer Observer) {
	observer.OnNext(e.element)
}

type errorEmission struct {
	err error
}

func (e *errorEmission) notify(observer Observer) {
	observer.OnError(e.err)
}

type completeEmission struct{}

func (e *completeEmission) notify(observer Observer) {
	observer.OnComplete()
}

type EmitterBuilder interface {
	EmitNext(el string) EmitterBuilder
	EmitError(err error) EmitterBuilder
	EmitComplete() EmitterBuilder
	Build() EmitSequence
}
type emitterBuilder struct {
	emissions []notifiable
}

func NewEmitterBuilder() EmitterBuilder {
	return &emitterBuilder{}
}

func (b *emitterBuilder) EmitNext(el string) EmitterBuilder {
	b.emissions = append(b.emissions, &nextEmission{element: el})
	return b
}

func (b *emitterBuilder) EmitError(err error) EmitterBuilder {
	b.emissions = append(b.emissions, &errorEmission{err: err})
	return b
}

func (b *emitterBuilder) EmitComplete() EmitterBuilder {
	b.emissions = append(b.emissions, &completeEmission{})
	return b
}

func (b *emitterBuilder) Build() EmitSequence {
	return func(observer Observer, state DisposableState) {
		for _, emission := range b.emissions {
			emission.notify(observer)
		}
	}
}
