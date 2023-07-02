package mistakes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/100_mistakes"
)

var _ = Describe("002UnnecessaryNestedCode", func() {

	Describe("first join implementation", func() {
		var (
			first, second string
			max uint
		)

		BeforeEach(func() {
			first, second = "first", "second"
			max = 15
		})

		When("s1 is empty", func() {
			BeforeEach(func() {
				first = ""
			})

			It("returns error because s1 is empty", func() {
				Ω(mistakes.Join(first, second, max)).Error().Should(MatchError(mistakes.ErrEmptyStringS1))
			})
		})

		When("s2 is empty", func() {
			BeforeEach(func() {
				second = ""
			})

			It("returns error because s2 is empty", func() {
				Ω(mistakes.Join(first, second, max)).Error().Should(MatchError(mistakes.ErrEmptyStringS2))
			})
		})

		When("everything goes fine", func() {
			It("returns the concatenation of the two string values", func() {
				Ω(mistakes.Join(first, second, max)).Should(Equal("firstsecond"))
			})
		})

		When("everything goes fine, but we need to truncate the response", func() {
			BeforeEach(func() {
				max = 7
			})

			It("returns the concatenation of the two string values", func() {
				Ω(mistakes.Join(first, second, max)).Should(Equal("firstse"))
			})
		})
	})

	Describe("uglyjoin implementation", func() {
		var (
			first, second string
			max uint
		)

		BeforeEach(func() {
			first, second = "first", "second"
			max = 15
		})

		When("s1 is empty", func() {
			BeforeEach(func() {
				first = ""
			})

			It("returns error because s1 is empty", func() {
				Ω(mistakes.UglyJoin(first, second, max)).Error().Should(MatchError(mistakes.ErrEmptyStringS1))
			})
		})

		When("s2 is empty", func() {
			BeforeEach(func() {
				second = ""
			})

			It("returns error because s2 is empty", func() {
				Ω(mistakes.UglyJoin(first, second, max)).Error().Should(MatchError(mistakes.ErrEmptyStringS2))
			})
		})

		When("everything goes fine", func() {
			It("returns the concatenation of the two string values", func() {
				Ω(mistakes.UglyJoin(first, second, max)).Should(Equal("firstsecond"))
			})
		})

		When("everything goes fine, but we need to truncate the response", func() {
			BeforeEach(func() {
				max = 7
			})

			It("returns the concatenation of the two string values", func() {
				Ω(mistakes.UglyJoin(first, second, max)).Should(Equal("firstse"))
			})
		})
	})
})
