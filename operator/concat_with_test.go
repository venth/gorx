package operator

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"
)

func TestOperator_ConcatWith(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operator.ConcatWith suite")
}

var _ = Describe("Operator.ConcatWith", func() {
	var (
		mockCtrl         *gomock.Controller
		emissionObserver *gorx.MockObserver
	)

	mockCtrl = gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	Context("when an observable is passed as the argument ", func() {
		It("concatenates the passed observable with the existing one", func() {
			elementFromExisting := "element from existing"
			elementFromConcatenated := "element from concatenated"
			gomock.InOrder(
				emissionObserver.EXPECT().OnNext(elementFromExisting).Times(1),
				emissionObserver.EXPECT().OnNext(elementFromConcatenated).Times(1),
				emissionObserver.EXPECT().OnComplete().Times(1),
			)

			Just(elementFromExisting).ConcatWith(Just(elementFromConcatenated)).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})

	})

})
