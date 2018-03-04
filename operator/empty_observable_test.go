package operator

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"

	"testing"
)

func TestObservable_Empty(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable.Empty suite")
}

var _ = Describe("Observable.Empty", func() {
	var (
		emptyObservable  gorx.Observable
		emissionObserver *gorx.MockObserver
	)

	mockCtrl := gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	It("emits complete", func() {
		emissionObserver.EXPECT().OnComplete().Times(1)

		emptyObservable.Subscribe(emissionObserver)
	})

	It("doesn't emit any element", func() {
		emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
		emissionObserver.EXPECT().OnComplete().Times(1)

		emptyObservable.Subscribe(emissionObserver)
	})

	It("doesn't emit error", func() {
		emissionObserver.EXPECT().OnError(gomock.Any()).Times(0)
		emissionObserver.EXPECT().OnComplete().Times(1)

		emptyObservable.Subscribe(emissionObserver)
	})

	BeforeEach(func() {
		emptyObservable = Empty()
		emissionObserver = gorx.NewMockObserver(mockCtrl)
	})
})
