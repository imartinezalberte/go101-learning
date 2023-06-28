package chapter1_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/exp/constraints"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter1"
)

// from - to range -> abs -> filter by even numbers -> sum the result
func imperativeStyle[T constraints.Integer](from, to T) T {
	if from == to {
		return 0
	}

	if from > to {
		from, to = to, from
	}

	var sum T
	for i := from; i <= to; i++ {
		if v := T(math.Abs(float64(i))); v % 2 == 0 {
			sum += v
		}
	}

	return sum
}

var _ = Describe("Stream", func() {

	Describe("using NumberRange as described in the book", func() {
		When("you do things the right manner", func() {
			var (
				from, to, output int
				predicate        chapter1.Predicate[int]
			)

			BeforeEach(func() {
				from, to, output = -5, 5, 30

				predicate = func(input int) bool {
					return input < 10
				}
			})

			It("is equal to the expected array", func() {
				Expect(chapter1.NumberRange(from, to).ToArr()).Should(HaveExactElements([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}))
			})

			It("is equal to the expected sum", func() {
				Expect(chapter1.NumberRange(from, to).Abs().Filter(predicate).Sum()).Should(Equal(output))
			})
		})

		When("Functional programming to create an integer range, use math.abs of it, filter them by even and sum the result", func() {
			var (
				from, to, output int
				predicate chapter1.Predicate[int]
			)

			BeforeEach(func() {
				from, to, output = -10, 10, 60

				predicate = func(input int) bool {
					return input % 2 == 0
				}
			})

			It("functional programming", func() {
				Expect(chapter1.NumberRange(from, to).Abs().Filter(predicate).Sum()).Should(Equal(output))
			})

			It("Imperative programming", func() {
				Expect(imperativeStyle(from, to)).Should(Equal(output))
			})
		})
	})
})
