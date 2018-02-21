package observer

import "github.com/venth/gorx"

type testObserver struct{
	elements []interface{}
	error error
	completed int
	disposable gorx.Disposable
}

type TestObserver interface {
	gorx.UnboundObserver
	ElementsCount() int
	Completed() bool
	HasError() bool
}

func NewTestObserver() TestObserver {
	return &testObserver{}
}

func (t *testObserver) ElementsCount() int {
	return len(t.elements)
}

func (t *testObserver) Completed() bool {
	return t.completed == 1
}

func (t *testObserver) HasError() bool {
	return t.error != nil
}

func (t *testObserver) Bind(disposable gorx.Disposable) gorx.BoundObserver {
	t.disposable = disposable
	//XXX correct
	return nil
}

func (t *testObserver) OnNext(element interface{}) {
	t.elements = append(t.elements, element)
}

func (t *testObserver) OnError(err error) {
	t.error = err
}

func (t *testObserver) OnComplete() {
	t.completed += 1
}
