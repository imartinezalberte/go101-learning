package chapter2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter2"
)

var _ = Describe("MapDispatcher", func() {
	Describe("Corner cases", func() {
		var operation string

		When("the operation don't exists", func() {
			BeforeEach(func() {
				operation = "nonexistent"
			})

			It("returns operation not found", func() {
				Expect(chapter2.Dispatch(operation, 0, 0)).Error().Should(HaveOccurred())
			})
		})

		When("the division operation is used with 0 as dividend", func() {
			BeforeEach(func() {
				operation = chapter2.DivisionOperation
			})

			It("returns 0 as default", func() {
				Expect(chapter2.Dispatch(operation, 10, 0)).Should(BeZero())
			})
		})
	})

	Describe("Using addition", func() {
		var operation string

		BeforeEach(func() {
			operation = chapter2.AdditionOperation
		})	

		When("the input is std int8", func() {
			var a, b, output int8

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint8", func() {
			var a, b, output uint8

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int16", func() {
			var a, b, output int16

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint16", func() {
			var a, b, output uint16

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int32", func() {
			var a, b, output int32

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint32", func() {
			var a, b, output uint32

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int64", func() {
			var a, b, output int64

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint64", func() {
			var a, b, output uint64

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int", func() {
			var a, b, output int

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint", func() {
			var a, b, output uint

			BeforeEach(func() {
				a, b, output = 2, 4, 6
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})
	})

	Describe("Using substraction", func() {
		var operation string

		BeforeEach(func() {
			operation = chapter2.SubstractionOperation
		})	

		When("the input is std int8", func() {
			var a, b, output int8

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint8", func() {
			var a, b, output uint8

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int16", func() {
			var a, b, output int16

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint16", func() {
			var a, b, output uint16

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int32", func() {
			var a, b, output int32

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint32", func() {
			var a, b, output uint32

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int64", func() {
			var a, b, output int64

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint64", func() {
			var a, b, output uint64

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std int", func() {
			var a, b, output int

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})

		When("the input is std uint", func() {
			var a, b, output uint

			BeforeEach(func() {
				a, b, output =  4, 2, 2
			})

			It("is the value expected", func() {
				Expect(chapter2.Dispatch(operation, a, b)).Should(Equal(output))
			})
		})
	})
})
