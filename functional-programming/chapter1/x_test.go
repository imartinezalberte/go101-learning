package chapter1_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter1"
)

var _ = Describe("X", func() {
	
	Describe("filter function", func() {
		When("Predicate of int", func() {
			var predicate chapter1.Predicate[int]

			BeforeEach(func() {
				predicate = func(input int) bool {
					return input % 2 == 0
				}
			})

			It("Contains exactly the even elements on the same order as input", func() {
				Expect(chapter1.Filter([]int{1, 2, 3, 4, 5, 6}, predicate)).Should(HaveExactElements(2, 4, 6))
			})
		})
	})
})
