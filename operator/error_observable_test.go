package operator

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"
)

func TestObservable_Error(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable.Error suite")
}

var _ = Describe("Observable.Error", func() {
	var (
		someError        error
		errorObservable  gorx.Observable
		emissionObserver *gorx.MockObserver
	)

	mockCtrl := gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	It("doesn't emit complete", func() {
		emissionObserver.EXPECT().OnComplete().Times(0)
		emissionObserver.EXPECT().OnError(gomock.Any()).Times(1)

		errorObservable.Subscribe(emissionObserver)
	})

	It("doesn't emit any element", func() {
		emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
		emissionObserver.EXPECT().OnError(gomock.Any()).Times(1)

		errorObservable.Subscribe(emissionObserver)
	})

	It("emits error", func() {
		emissionObserver.EXPECT().OnError(someError).Times(1)

		errorObservable.Subscribe(emissionObserver)
	})

	BeforeEach(func() {
		someError = errors.New("some error")
		errorObservable = Error(someError)
		emissionObserver = gorx.NewMockObserver(mockCtrl)
	})
})
