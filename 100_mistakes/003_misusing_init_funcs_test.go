package mistakes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/100_mistakes"
)

var _ = Describe("003MisusingInitFuncs", func() {
	Describe("DBConn global variable is harder to test", func() {
		When("you don't want to modify the global variable, everything is fine", func() {
			It("", func() {
				Ω(mistakes.GlobalDBConn.Execute()).Error().ShouldNot(HaveOccurred())
			})
		})

		When("how to modify the global variable?", func() {
			It("is impossible right now", func() {
			})
		})
	})
})
