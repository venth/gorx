package subscription

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"
)


func TestQueuingSubscription(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queuing Subscription Suite")
}

var _ = Describe("Queuing Subscription", func() {
	mockCtrl := gomock.NewController(GinkgoT())
	defer mockCtrl.Finish()

	Context("which is newly created", func() {
		var (
			subscription gorx.Subscription
			unboundObserver *gorx.MockUnboundObserver
			boundObserver *gorx.MockBoundObserver
			emitter gorx.Emitter
		)
		BeforeEach(func() {
			boundObserver = gorx.NewMockBoundObserver(mockCtrl)
			unboundObserver = gorx.NewMockUnboundObserver(mockCtrl)
			unboundObserver.EXPECT().Bind(gomock.Any()).Return(boundObserver)
			boundObserver.EXPECT().Unbind().Return(unboundObserver)
		})

		It("delegates complete to subscribed enqueuingObserver", func() {
			emitter = gorx.NewEmitterBuilder().EmitComplete().Build()
			boundObserver.EXPECT().OnComplete().Times(1)

			subscription = NewQueuingSubscription(emitter, unboundObserver)

		})

		It("delegates next to subscribed enqueuingObserver", func() {
			element := "first element"
			emitter = gorx.NewEmitterBuilder().EmitNext(element).Build()
			boundObserver.EXPECT().OnNext(element).Times(1)

			subscription = NewQueuingSubscription(emitter, unboundObserver)

		})

		It("delegates all next to subscribed enqueuingObserver", func() {
			firstElement := "first firstElement"
			secondElement := "second firstElement"
			emitter = gorx.NewEmitterBuilder().
				EmitNext(firstElement).
				EmitNext(secondElement).
				Build()
			boundObserver.EXPECT().OnNext(firstElement).Times(1)
			boundObserver.EXPECT().OnNext(secondElement).Times(1)

			subscription = NewQueuingSubscription(emitter, unboundObserver)
		})

		It("delegates error to subscribed enqueuingObserver", func() {
			err := errors.New("an error")
			emitter = gorx.NewEmitterBuilder().EmitError(err).Build()
			boundObserver.EXPECT().OnError(err).Times(1)

			subscription = NewQueuingSubscription(emitter, unboundObserver)
		})

		It("stops to handle anything when complete is received", func() {
			element := "element"
			emitter = gorx.NewEmitterBuilder().
				EmitComplete().
				EmitNext(element).
				Build()

			boundObserver.EXPECT().OnComplete().Times(1)
			boundObserver.EXPECT().OnNext(gomock.Any()).Times(0)

			subscription = NewQueuingSubscription(emitter, unboundObserver)
		})
	})
})

