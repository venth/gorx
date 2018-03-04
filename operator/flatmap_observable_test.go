package operator

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"

	"testing"
)

func TestObservable_FlatMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable.FlatMap suite")
}

var _ = PDescribe("Observable.FlatMap", func() {

	var emissionObserver *gorx.MockObserver
	emptyObservable := CreateObservable(func(emitter gorx.Observer, state gorx.DisposableState) {
		emitter.OnComplete()
	})

	Context("when empty observable emits no elements", func() {
		emptyObservable = emptyObservable.FlatMap(func(el interface{}) gorx.Observable {
			return CreateObservable(func(emitter gorx.Observer, state gorx.DisposableState) {
				emitter.OnNext(el)
				emitter.OnComplete()
			})
		})

		It("returns empty observable", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			emptyObservable.Subscribe(emissionObserver)
		})

	})

	mockCtrl := gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	BeforeEach(func() {
		emissionObserver = gorx.NewMockObserver(mockCtrl)
	})
})
