package pointers_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/pointers"
)

var _ = Describe("Pointers", func() {
	When("creating pointers using new function", func() {
		var ptr *pointers.Ptr

		BeforeEach(func() {
			ptr = new(pointers.Ptr) // *pointers.Ptr
		})

		It("value returned by new is not nil", func() {
			Ω(ptr).NotTo(BeNil()) // Pointer is not nil
		})

		It("value returned by the pointer is a nil value (because that's the variable's default)", func() {
			Ω(*ptr).To(BeNil()) // Where the pointer points, obviously is a nil value
		})

		It("type of ptr is *pointers.Ptr", func() {
			Ω(fmt.Sprintf("%T", ptr)).To(Equal("*pointers.Ptr"))
		})

		It("type of *ptr is pointers.Ptr", func() {
			Ω(fmt.Sprintf("%T", *ptr)).To(Equal("pointers.Ptr"))
		})
	})

	When("comparing two pointers", func() {
		var (
			i          *int
			tmp1, tmp2 *int
		)

		BeforeEach(func() {
			i = new(int)
			tmp1, tmp2 = i, i
		})

		It("results in two pointers that are equal because they have the same address", func() {
			Ω(tmp1).To(Equal(tmp2))
		})

		It("contains the same address", func() {
			Ω(fmt.Sprint(tmp1)).To(Equal(fmt.Sprint(tmp2)))
		})
	})

	When("a pointer can't be converted to an arbitrary type", func() {
		var (
			i     *int64
			myint *pointers.MyInt
		)

		BeforeEach(func() {
			i = new(int64)
			myint = new(pointers.MyInt)
		})

		It("values of type *int64 can be implicitly converted to type pointers.Ta", func() {
			var x pointers.Ta

			x = i

			_ = x
		})

		It("values of type pointers.Ta can be implicitly converted to type *int64", func() {
			var x *int64 = *new(pointers.Ta)

			_ = x
		})

		It("values of type *pointers.MyInt can be implicitly converted to type pointers.Tb", func() {
			var x pointers.Tb

			x = myint

			_ = x
		})

		It("values of type pointers.Tb can be implicitly converted to type *pointers.MyInt", func() {
			var x *pointers.MyInt = *new(pointers.Tb)

			_ = x
		})

		It("values of type *pointers.MyInt can be explicitly converted to type *int64", func() {
			var x *pointers.MyInt = (*pointers.MyInt)(i)

			_ = x
		})

		It("values of type *int64 can be explicitly converted to type *pointers.MyInt", func() {
			var x *int64 = (*int64)(myint)

			_ = x
		})

		It("values of type pointers.Ta can be converted to pointers.Tb", func() {
			var (
				ta pointers.Ta = i
				tb pointers.Tb = pointers.Tb((*pointers.MyInt)((*int64)(ta)))
			)

			_ = tb
		})
	})
})
