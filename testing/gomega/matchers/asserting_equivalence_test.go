package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type (
	i8  int8
	u8  uint8
	i16 int16
	u16 uint16
	i32 int32
	u32 uint32
	i64 int64
	u64 uint64
)

var _ = Describe("AssertingEquivalence", func() {
	Describe("Equal method", func() {
		DescribeTable("comparing primitives", func(input, expected any) {
			Ω(input).Should(Equal(expected))
		},
			Entry("int8s", i8(1), i8(1)),
			Entry("uint8s", u8(1), u8(1)),
			Entry("int16s", i16(1), i16(1)),
			Entry("uint16s", u16(1), u16(1)),
			Entry("int32s", i32(1), i32(1)),
			Entry("uint32s", u32(1), u32(1)),
			Entry("int64s", i64(1), i64(1)),
			Entry("uint64s", u64(1), u64(1)),
			Entry("ints", int(1), int(1)),
			Entry("uints", uint(1), uint(1)),
			// Entry("ints using types", int8(1), i8(1)), // Will fail because they are not the same type
			Entry("uintptrs", uintptr(1), uintptr(1)),
			Entry("float32s", float32(1), float32(1)),
			Entry("float64s", float64(1), float64(1)),
			Entry("strings", "hello", "hello"),
		)

		When("Interesting ginkgo cases", func() {
			// We are using Pending It, because it will always fail
			PIt("refusing to compare nil", func() {
				Ω(nil).Should(Equal(nil))
			})

			It("comparing bytes is using bytes.Equal instead of reflect.DeepEqual", func() {
				Ω([]byte("hello")).Should(Equal([]byte("hello")))
			})
		})
	})

	Describe("BeEquivalentTo", func() {
		DescribeTable("comparing primitives", func(input, expected any) {
			Ω(input).Should(BeEquivalentTo(expected))
		},
			Entry("int8s", i8(1), int8(1)),
			Entry("uint8s", u8(1), uint8(1)),
			Entry("int16s", i16(1), int16(1)),
			Entry("uint16s", u16(1), uint16(1)),
			Entry("int32s", i32(1), int32(1)),
			Entry("uint32s", u32(1), uint32(1)),
			Entry("int64s", i64(1), int64(1)),
			Entry("uint64s", u64(1), uint64(1)),
		)
	})

	Describe("BeNumerically", func() {
		DescribeTable("comparing primitives", func(input, expected any) {
			Ω(input).Should(BeNumerically(expected))
		},
			Entry("int8s", i8(1), int8(1)),
			Entry("uint8s", u8(1), uint8(1)),
			Entry("int16s", i16(1), int16(1)),
			Entry("uint16s", u16(1), uint16(1)),
			Entry("int32s", i32(1), int32(1)),
			Entry("uint32s", u32(1), uint32(1)),
			Entry("int64s", i64(1), int64(1)),
			Entry("uint64s", u64(1), uint64(1)),
		)
	})
})
