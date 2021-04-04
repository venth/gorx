package operator

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/venth/gorx"
)

var _ = Describe("Observable.Concat", func() {
	var (
		emissionObserver *gorx.MockObserver
		mockCtrl         *gomock.Controller
	)

	mockCtrl = gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	someElement := "some element"

	Context("of empty sequences", func() {
		It("emits complete for one empty sequence", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			Concat(Empty()).Subscribe(emissionObserver)
		})

		It("emits complete for two empty sequences", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			Concat(Empty(), Empty()).Subscribe(emissionObserver)
		})

		It("emits complete for three empty sequences", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			Concat(Empty(), Empty(), Empty()).Subscribe(emissionObserver)
		})

		It("emits elements for non empty sequences", func() {
			gomock.InOrder(
				emissionObserver.EXPECT().OnNext(someElement).Times(1),
				emissionObserver.EXPECT().OnComplete().Times(1),
			)

			Concat(Empty(), Just(someElement), Empty()).Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("without observables to concatenate", func() {
		It("emits complete", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			Concat().Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("of error sequences", func() {
		It("emits elements till error is occurred", func() {
			someError := errors.New("some error")

			gomock.InOrder(
				emissionObserver.EXPECT().OnNext(someElement).Times(1),
				emissionObserver.EXPECT().OnError(someError).Times(1),
			)

			Concat(Just(someElement), Error(someError), Just(someElement)).Subscribe(emissionObserver)
		})

		It("stops on first error", func() {
			someError := errors.New("some error")

			gomock.InOrder(
				emissionObserver.EXPECT().OnError(someError).Times(1),
			)

			Concat(Error(someError), Error(someError), Just(someElement)).Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

	Context("of sequences", func() {
		It("emits elements in order", func() {
			element1 := "element1"
			element2 := "element2"
			element3 := "element3"
			element4 := "element4"
			element5 := "element5"
			element6 := "element6"

			gomock.InOrder(
				emissionObserver.EXPECT().OnNext(element1).Times(1),
				emissionObserver.EXPECT().OnNext(element2).Times(1),
				emissionObserver.EXPECT().OnNext(element3).Times(1),
				emissionObserver.EXPECT().OnNext(element4).Times(1),
				emissionObserver.EXPECT().OnNext(element5).Times(1),
				emissionObserver.EXPECT().OnNext(element6).Times(1),
				emissionObserver.EXPECT().OnComplete().Times(1),
			)

			Concat(
				Just(element1),
				Just(element2),
				Just(element3),
				Just(element4),
				Just(element5),
				Just(element6),
			).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})
})
