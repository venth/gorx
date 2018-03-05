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
		mockCtrl *gomock.Controller
	)

	mockCtrl = gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	Context("when emitter didn't emit any error", func() {
		It("emits original elements", func() {
			emissionObserver.EXPECT().OnNext("element").Times(1)
			emissionObserver.EXPECT().OnComplete().Times(1)

			Just("element").OnErrorResumeNext(Just("resumed")).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("when emitter emits only complete", func() {
		It("emits only complete as well", func() {
			emissionObserver.EXPECT().OnComplete().Times(1)

			Empty().OnErrorResumeNext(Just("resumed")).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

	Context("when emitter an error", func() {
		It("emits resumed", func() {
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnNext("resumed").Times(1)
			emissionObserver.EXPECT().OnComplete().Times(1)

			Error(errors.New("some error")).
				OnErrorResumeNext(Just("resumed")).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})

})
