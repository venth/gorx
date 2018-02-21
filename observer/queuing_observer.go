package observer

import "github.com/venth/gorx"

func NewQueuingObserver(emitter gorx.Emitter, targetObserver gorx.UnboundObserver) gorx.UnboundObserver {
	return &queuingObserver{
		emitter: emitter,
		targetObserver: targetObserver,
	}
}

type commandType int

const (
	CtNext     commandType = iota
	CtErr
	CtComplete
	CtDispose
)

type emittedCommand struct {
	emittedCommandType commandType
	element interface{}
}

var executeCommand = map[commandType] func(interface{}, gorx.Observer) {
	CtNext: func(element interface{}, observer gorx.Observer) {
		observer.OnNext(element)
	},
	CtErr: func(err interface{}, observer gorx.Observer) {
		observer.OnError(err.(error))
	},
	CtComplete: func(irrelevant interface{}, observer gorx.Observer) {
		observer.OnComplete()
	},
	CtDispose: func(irrelevant interface{}, observer gorx.Observer) {

	},
}


type queuingObserver struct {
	emitter        gorx.Emitter
	targetObserver gorx.UnboundObserver
}

type boundQueuingObserver struct {
	done chan struct{}
	commands       chan *emittedCommand
	emitter        gorx.Emitter
	targetObserver gorx.BoundObserver
}

func (o *queuingObserver) OnNext(element interface{}) {
}

func (o *queuingObserver) OnError(err error) {
}

func (o *queuingObserver) OnComplete() {
}

func (o *queuingObserver) Bind(disposable gorx.Disposable) gorx.BoundObserver {
	bound := &boundQueuingObserver{
		emitter: o.emitter,
		targetObserver: o.targetObserver.Bind(disposable),
		commands: make(chan *emittedCommand),
		done: make(chan struct{}),
	}

	go bound.Run()

	return bound
}

func (o *boundQueuingObserver) Run() {
	defer close(o.commands)
	go o.emitter(o.targetObserver)
	o.handleEmissionQueue()
}

func (o *boundQueuingObserver) handleEmissionQueue() {
	for {
		cmd, more := <-o.commands
		if more {
			executeCommand[cmd.emittedCommandType](cmd.element, o.targetObserver)
		} else {
			break
		}
	}
}

func (o *boundQueuingObserver) OnNext(element interface{}) {
	o.commands <- &emittedCommand{emittedCommandType: CtNext, element:element}
}

func (o *boundQueuingObserver) OnError(err error) {
	o.commands <- &emittedCommand{emittedCommandType: CtErr, element:err}
}

func (o *boundQueuingObserver) OnComplete() {
	o.commands <- &emittedCommand{emittedCommandType: CtComplete}
}


func (o *boundQueuingObserver) Unbind() gorx.UnboundObserver {
	unboundTargetObserver := o.targetObserver.Unbind()
	close(o.done)

	return &queuingObserver{emitter: o.emitter, targetObserver: unboundTargetObserver}
}
