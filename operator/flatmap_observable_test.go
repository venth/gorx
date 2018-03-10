package operator

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"
)

func TestObservable_FlatMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable.FlatMap suite")
}

var _ = Describe("Observable.FlatMap", func() {

	var emissionObserver *gorx.MockObserver
	someElement := "some element"

	mockCtrl := gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	Context("when empty observable emits no elements", func() {
		emptyObservable := Empty()

		It("returns empty observable", func() {
			emptyObservable.FlatMap(func(el interface{}) gorx.Observable {
				return Just("element")
			})
			emissionObserver.EXPECT().OnComplete().Times(1)

			emptyObservable.Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

	Context("when observable emits elements", func() {
		It("flattens them", func() {
			flatten := Just(someElement, someElement).
				FlatMap(func(el interface{}) gorx.Observable {
				return Just(el)
			})

			emissionObserver.EXPECT().OnNext(someElement).Times(2)
			emissionObserver.EXPECT().OnComplete().Times(1)

			flatten.Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

	Context("when observable emits slices", func() {
		slice := []string{someElement, someElement, someElement}
		Context("one slice", func() {
			observable := Just(slice)
			It("stops to flatten them because of an error after first element", func() {
				someError := fmt.Errorf("some error")

				flatten := observable.FlatMap(func(el interface{}) gorx.Observable {
					return Just(someElement).ConcatWith(Error(someError)).ConcatWith(FromSlice(el))
				})

				emissionObserver.EXPECT().OnNext(someElement).Times(1)
				emissionObserver.EXPECT().OnError(someError).Times(1)

				flatten.Subscribe(emissionObserver)
			})
		})
		Context("two slices", func() {
			observable := Just(slice, slice)
			It("flattens them with an error after first element", func() {
				someError := fmt.Errorf("some error")

				flatten := observable.FlatMap(func(el interface{}) gorx.Observable {
					return Just(someElement).ConcatWith(Error(someError)).ConcatWith(FromSlice(el))
				})

				emissionObserver.EXPECT().OnNext(someElement).Times(1)
				emissionObserver.EXPECT().OnError(someError).Times(1)

				flatten.Subscribe(emissionObserver)
			})
		})


		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})
})
