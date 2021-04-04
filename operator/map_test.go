package operator

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/venth/gorx"
)

var _ = Describe("Observable.Map", func() {
	mockCtrl := gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	var emissionObserver *gorx.MockObserver

	Context("when emits one element", func() {
		oneElement := "one element"
		mappedElement := "mapped element"
		emission := gorx.NewEmitterBuilder().
			EmitNext(oneElement).
			EmitComplete().
			Build()

		observable := CreateObservable(emission)

		It("maps the element to mapped element", func() {
			emissionObserver.EXPECT().OnNext(mappedElement + oneElement)
			emissionObserver.EXPECT().OnComplete()

			observable.Map(func(el interface{}) interface{} {
				return mappedElement + el.(string)
			}).
				Subscribe(emissionObserver)
		})

		BeforeEach(func() {
			emissionObserver = gorx.NewMockObserver(mockCtrl)
		})
	})
})
