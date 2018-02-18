package observable

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/venth/gorx"
	"github.com/venth/gorx/observer"
)

func TestObservable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable Suite")
}

var _ = Describe("Observable", func() {
	var ob gorx.Observable

	Context("a fresh one", func() {
		It("creates empty observable", func() {
			ob = Empty()

			Expect(ob).Should(Not(BeNil()))
		})
	})

	Context("with subscribed observer", func() {
		It("returns subscription after observer subscribe itself", func() {
			ob = Empty()
			someObserver := observer.TestObserver()

			subscription := ob.Subscribe(someObserver)

			Expect(subscription).Should(Not(BeNil()))
		})

		It("returns not disposed subscription", func() {
			ob = Empty()
			someObserver := observer.TestObserver()

			subscription := ob.Subscribe(someObserver)

			Expect(subscription.IsDisposed()).Should(BeFalse())
		})

		It("disposes subscription", func() {
			ob = Empty()
			someObserver := observer.TestObserver()
			subscription := ob.Subscribe(someObserver)
			disposedSubscription := subscription.Dispose()

			Expect(disposedSubscription.IsDisposed()).Should(BeTrue())
		})

	})

})
