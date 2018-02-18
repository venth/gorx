package observer

import "github.com/venth/gorx"

type testObserver struct{
	gorx.Observer
}

func TestObserver() gorx.Observer {
	return &testObserver{}
}
