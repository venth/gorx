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

	Context("which is a fresh one", func() {
		It("creates empty observable", func() {
			ob = Empty()

			Expect(ob).Should(Not(BeNil()))
		})
	})

	Context("which is an empty one", func() {

		It("emits no elements", func() {
			empty := Empty()
			observer := observer.NewTestObserver()

			empty.Subscribe(observer)

			Expect(observer.ElementsCount()).Should(BeZero())
		})
		It("emits no errors", func() {
			empty := Empty()
			observer := observer.NewTestObserver()

			empty.Subscribe(observer)

			Expect(observer.HasError()).Should(BeFalse())
		})
		It("completes after subscription", func() {
			empty := Empty()
			observer := observer.NewTestObserver()

			empty.Subscribe(observer)

			Expect(observer.Completed()).Should(BeTrue())
		})
	})

	Context("with subscribed observer", func() {
		It("returns subscription after observer subscribe itself", func() {
			ob = Empty()
			someObserver := observer.NewTestObserver()

			subscription := ob.Subscribe(someObserver)

			Expect(subscription).Should(Not(BeNil()))
		})

		It("returns not disposed subscription", func() {
			ob = Empty()
			someObserver := observer.NewTestObserver()

			subscription := ob.Subscribe(someObserver)

			Expect(subscription.IsDisposed()).Should(BeFalse())
		})

		It("disposes subscription", func() {
			ob = Empty()
			someObserver := observer.NewTestObserver()
			subscription := ob.Subscribe(someObserver)
			disposedSubscription := subscription.Dispose()

			Expect(disposedSubscription.IsDisposed()).Should(BeTrue())
		})

	})

})
