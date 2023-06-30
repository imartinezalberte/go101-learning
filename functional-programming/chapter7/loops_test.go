package chapter7_test

import (
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter7"
)

var _ = Describe("Loops", func() {
	Describe("StdLoops", func() {
		var (
			from, to, steps int
			accept          chapter7.Consumer[int]
		)

		When("Everything is fine", func() {
			var out strings.Builder

			BeforeEach(func() {
				from, to, steps = 1, 10, 1
				accept = func(input int) {
					out.WriteString(strconv.Itoa(input))
				}
			})

			It("We should receive our number list from 'from' to 'to' following 'steps' value", func() {
				chapter7.StdLoop(from, to, steps, accept)
				Ω(out.String()).Should(Equal("123456789"))
			})
		})

		When("From is greater than to", func() {
			var out strings.Builder

			BeforeEach(func() {
				from, to, steps = 1, 0, 1
				accept = func(input int) {
					out.WriteString(strconv.Itoa(input))
				}
			})

			It("We should receive nothing", func() {
				chapter7.StdLoop(from, to, steps, accept)
				Ω(out.String()).Should(BeEmpty())
			})
		})

		When("from is greater than to and steps is positive", func() {
			var out strings.Builder
			BeforeEach(func() {
				from, to, steps = 10, 1, 1
				accept = func(input int) {
					out.WriteString(strconv.Itoa(input))
				}
			})

			It("We should receive nothing", func() {
				Ω(out.String()).Should(BeEmpty())
			})
		})

		When("from is greater than to and steps is negative", func() {
			var out strings.Builder
			BeforeEach(func() {
				from, to, steps = 10, 1, -1
				accept = func(input int) {
					out.WriteString(strconv.Itoa(input))
				}
			})

			It("We should receive 10987654321", func() {
				Ω(out.String()).Should(Equal("10987654321"))
			})
		})
	})
})
