package operator

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"

	"testing"
)

func TestObservable_Just(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable.Just suite")
}

var _ = Describe("Observable.Just", func() {
	var (
		elements         []interface{}
		justObservable   gorx.Observable
		emissionObserver *gorx.MockObserver
		mockCtrl         *gomock.Controller
	)

	Context("when there are no elements", func() {
		justObservable = Just()

		It("completes without emitting any element", func() {
			emissionObserver.EXPECT().OnNext(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnError(gomock.Any()).Times(0)
			emissionObserver.EXPECT().OnComplete().Times(1)

			justObservable.Subscribe(emissionObserver)
		})
	})

	mockCtrl = gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	BeforeEach(func() {
		elements = make([]interface{}, 0)
		emissionObserver = gorx.NewMockObserver(mockCtrl)
	})
})
