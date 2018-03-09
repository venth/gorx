package operator

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"
)

func TestOperator_OnErrorResumeNext(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operator.OnErrorResumeNext suite")
}

var _ = Describe("Operator.OnErrorResumeNext", func() {
	var (
		emissionObserver *gorx.MockObserver
		mockCtrl         *gomock.Controller
	)

	mockCtrl = gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	someResumed := "some resumed"
	someResumedFunc := func(error) gorx.Observable { return Just(someResumed) }

	someError := errors.New("some error")

	Context("when emitter didn't emit any error", func() {

		It("emits original elements", func() {
			originalElement := "original element"

			gomock.InOrder(
				emissionObserver.EXPECT().OnNext(originalElement).Times(1),
				emissionObserver.EXPECT().OnComplete().Times(1),
			)

			Just(originalElement).OnErrorResumeNext(someResumedFunc).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("when emitter emits only complete", func() {
		It("emits only complete as well", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			Empty().OnErrorResumeNext(someResumedFunc).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("when emitter emits on error in middle of sequence", func() {
		It("emits resumed sequence and doesn't emit elements after error has occurred", func() {
			someElement := "some element"

			gomock.InOrder(
				emissionObserver.EXPECT().OnNext(someElement).Times(1),
				emissionObserver.EXPECT().OnNext(someResumed).Times(1),
				emissionObserver.EXPECT().OnComplete().Times(1),
			)

			Concat(
				Just(someElement),
				Error(someError),
				Just(someElement),
			).OnErrorResumeNext(someResumedFunc).
				Subscribe(emissionObserver)
		})
		
		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("when emitter an error", func() {
		It("emits resumed observed sequence, which doesn't contain errors", func() {
			resumedElement := "resumed"
			resumedFunc := func(error) gorx.Observable { return Just(resumedElement) }

			emissionObserver.EXPECT().OnNext(resumedElement).Times(1)
			emissionObserver.EXPECT().OnComplete().Times(1)

			Error(someError).
				OnErrorResumeNext(resumedFunc).
				Subscribe(emissionObserver)
		})

		It("emits resumed observed sequence, which contain error", func() {
			resumedWithError := errors.New("resumed error")
			resumedWithErrorFunc := func(error) gorx.Observable { return Error(resumedWithError) }

			emissionObserver.EXPECT().OnError(resumedWithError).Times(1)

			Error(someError).
				OnErrorResumeNext(resumedWithErrorFunc).
				Subscribe(emissionObserver)
		})

		It("completes because resumed sequence is empty", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			resumedWithEmptyFunc := func(error) gorx.Observable { return Empty() }
			Error(someError).
				OnErrorResumeNext(resumedWithEmptyFunc).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

})
