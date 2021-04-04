package operator

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/venth/gorx"
)

var _ = Describe("Observable.Just", func() {
	var (
		justObservable   gorx.Observable
		emissionObserver *gorx.MockObserver
		mockCtrl         *gomock.Controller
	)

	Context("when there are no elements", func() {

		It("completes without emitting any element", func() {
			justObservable = Just()

			emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnComplete().Times(1)

			justObservable.Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

	Context("when there are elements", func() {

		It("completes after just one element", func() {
			element := "element"
			justObservable = Just(element)

			emissionObserver.EXPECT().OnNext(element).Times(1)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnComplete().Times(1)

			justObservable.Subscribe(emissionObserver)
		})

		It("completes after just two elements", func() {
			firstElement := "firstElement"
			secondElement := "secondElement"
			justObservable = Just(firstElement, secondElement)

			observedFirstElement := emissionObserver.EXPECT().OnNext(firstElement).Times(1)
			emissionObserver.EXPECT().OnNext(secondElement).Times(1).After(observedFirstElement)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnComplete().Times(1)

			justObservable.Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

	Context("when passed elements are nil", func() {
		It("emits error for nil element", func() {
			justObservable = Just(nil)
			emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(1)
			emissionObserver.EXPECT().OnComplete().Times(0)

			justObservable.Subscribe(emissionObserver)
		})

		It("stops emission after first nil element", func() {
			justObservable = Just(nil, nil)
			emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(1)
			emissionObserver.EXPECT().OnComplete().Times(0)

			justObservable.Subscribe(emissionObserver)
		})

		It("stops emission after first nil element, event if the next one isn't nil", func() {
			justObservable = Just(nil, "some element")
			emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(1)
			emissionObserver.EXPECT().OnComplete().Times(0)

			justObservable.Subscribe(emissionObserver)
		})

		It("emits not nil element and stops emission on nil element", func() {
			justObservable = Just("some element", nil)
			emissionObserver.EXPECT().OnNext(gomock.Any()).Times(1)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(1)
			emissionObserver.EXPECT().OnComplete().Times(0)

			justObservable.Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

	mockCtrl = gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()
})
