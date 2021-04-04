package operator

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/venth/gorx"
)

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
