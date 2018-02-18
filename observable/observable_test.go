package observable

import (
	"testing"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestObservable(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "Observable Suite")
}


var _ = Describe("Observable", func() {
	Context("When observable isn't observed", func() {
		It("is observed after an observer subscribes to it", func() {

		})
	})

})