package chapter6_test

import (
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter6"
)

var _ = Describe("C", func() {
	Describe("DropWhile", func() {
		var (
			input []int
			pred  chapter6.Predicate[int]
		)

		When("There's something to drop", func() {
			var (
				predGenFn func(int) chapter6.Predicate[int]
				threshold int
			)

			BeforeEach(func() {
				input = []int{0, 1, 2, 3, 4, 5}
				predGenFn = func(threshold int) chapter6.Predicate[int] {
					return func(input int) bool {
						return input < threshold
					}
				}
			})

			When("not dropping nothing", func() {
				BeforeEach(func() {
					threshold = 0
					pred = predGenFn(threshold)
				})

				It("returns [0, 1, 2, 3, 4, 5]", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(HaveExactElements(input[threshold:]))
				})
			})

			When("dropping first element", func() {
				BeforeEach(func() {
					threshold = 1
					pred = predGenFn(threshold)
				})

				It("returns [1, 2, 3, 4, 5]", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(HaveExactElements(input[threshold:]))
				})
			})

			When("dropping til second element", func() {
				BeforeEach(func() {
					threshold = 2
					pred = predGenFn(threshold)
				})

				It("returns [2, 3, 4, 5]", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(HaveExactElements(input[threshold:]))
				})
			})

			When("dropping til third element", func() {
				BeforeEach(func() {
					threshold = 3
					pred = predGenFn(threshold)
				})

				It("returns [3, 4, 5]", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(HaveExactElements(input[threshold:]))
				})
			})

			When("dropping til fourth element", func() {
				BeforeEach(func() {
					threshold = 4
					pred = predGenFn(threshold)
				})

				It("returns [4, 5]", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(HaveExactElements(input[threshold:]))
				})
			})

			When("dropping til fifth element", func() {
				BeforeEach(func() {
					threshold = 5
					pred = predGenFn(threshold)
				})

				It("returns [5]", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(HaveExactElements(input[threshold:]))
				})
			})

			When("dropping everything", func() {
				BeforeEach(func() {
					threshold = 6
					pred = predGenFn(threshold)
				})

				It("returns []", func() {
					Expect(chapter6.DropWhile(input, pred)).Should(BeEmpty())
				})
			})
		})
	})

	Describe("TakeWhile", func() {
		var (
			input []int
			pred  chapter6.Predicate[int]
		)

		When("There's something to take", func() {
			var (
				predGenFn func(int) chapter6.Predicate[int]
				threshold int
			)

			BeforeEach(func() {
				input = []int{0, 1, 2, 3, 4, 5}
				predGenFn = func(threshold int) chapter6.Predicate[int] {
					return func(input int) bool {
						return input <= threshold
					}
				}
			})

			When("taking everything", func() {
				BeforeEach(func() {
					threshold = 5
					pred = predGenFn(threshold)
				})

				It("returns [0, 1, 2, 3, 4, 5]", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(HaveExactElements(input[:threshold+1]))
				})
			})

			When("taking form first element", func() {
				BeforeEach(func() {
					threshold = 4
					pred = predGenFn(threshold)
				})

				It("returns [0, 1, 2, 3, 4]", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(HaveExactElements(input[:threshold+1]))
				})
			})

			When("taking from second element", func() {
				BeforeEach(func() {
					threshold = 3
					pred = predGenFn(threshold)
				})

				It("returns [0, 1, 2, 3]", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(HaveExactElements(input[:threshold+1]))
				})
			})

			When("taking from third element", func() {
				BeforeEach(func() {
					threshold = 2
					pred = predGenFn(threshold)
				})

				It("returns [0, 1, 2]", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(HaveExactElements(input[:threshold+1]))
				})
			})

			When("taking from fourth element", func() {
				BeforeEach(func() {
					threshold = 1
					pred = predGenFn(threshold)
				})

				It("returns [0, 1]", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(HaveExactElements(input[:threshold+1]))
				})
			})

			When("taking from fifth element", func() {
				BeforeEach(func() {
					threshold = 0
					pred = predGenFn(threshold)
				})

				It("returns [0]", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(HaveExactElements(input[:threshold+1]))
				})
			})

			When("taking nothing", func() {
				BeforeEach(func() {
					threshold = -1
					pred = predGenFn(threshold)
				})

				It("returns []", func() {
					Expect(chapter6.TakeWhile(input, pred)).Should(BeEmpty())
				})
			})
		})
	})

	Describe("mapping persons by gender", func() {
		type (
			Gender string
			person struct {
				Name   string
				Gender Gender
			}
		)

		const (
			Male   Gender = "male"
			Female Gender = "female"
		)

		var persons []person

		When("mapping persons, output name should contain Mr and Mrs", func() {
			var id func(int, interface{}) string

			BeforeEach(func() {
				persons = []person{{
					Name:   "Guillermo",
					Gender: Male,
				}, {
					Name:   "Juana",
					Gender: Female,
				}}
				id = func(index int, _ interface{}) string {
					return strconv.Itoa(index)
				}
			})

			It("male -> Mr and female -> Mrs", func() {
				Expect(chapter6.Map(persons, func(p person) person {
					switch p.Gender {
					case Male:
						p.Name = "Mr " + p.Name
					default:
						p.Name = "Mrs " + p.Name
					}
					return p
				})).Should(MatchAllElementsWithIndex(id, Elements{
					"0": MatchFields(IgnoreExtras, Fields{
						"Name": Equal("Mr Guillermo"),
					}),
					"1": MatchFields(IgnoreExtras, Fields{
						"Name": Equal("Mrs Juana"),
					}),
				}))
			})
		})
	})

	Describe("reducing slices", func() {
		
		When("Sum a slice of ints", func() {
			var (
				input []int
				expected int
			)

			BeforeEach(func() {
				input = []int{1, 2, 3, 4}
				expected = 10
			})

			It("10", func() {
				Expect(chapter6.Reduce(input, func(first, second int) int {
					return first + second
				})).Should(Equal(expected))
			})

			It("10", func() {
				Expect(chapter6.ReduceLazy(input, func(first, second int) int {
					return first + second
				})).Should(Equal(expected))
			})
		})

		When("join strings", func() {
			var (
				input []string
				expected string
			)

			BeforeEach(func() {
				expected = "hello"
				input = strings.Split(expected, "")
			})

			It("hello", func() {
				Expect(chapter6.Reduce(input, func(first, second string) string {
					return first + second
				})).Should(Equal(expected))
			})

			It("hello", func() {
				Expect(chapter6.ReduceLazy(input, func(first, second string) string {
					return first + second
				})).Should(Equal(expected))
			})
		})
	})

	Describe("reducing slices with start", func() {
		
		When("Sum a slice of ints", func() {
			var (
				input []int
				start, expected int
			)

			BeforeEach(func() {
				input = []int{1, 2, 3, 4}
				start, expected = 5, 15
			})

			It("15", func() {
				Expect(chapter6.ReduceWithStart(input, start, func(first, second int) int {
					return first + second
				})).Should(Equal(expected))
			})
		})

		When("join strings", func() {
			var (
				input []string
				start, expected string
			)

			BeforeEach(func() {
				start, expected = "hello", " world"
				input = strings.Split(expected, "")
			})

			It("hello world", func() {
				Expect(chapter6.ReduceWithStart(input, start, func(first, second string) string {
					return first + second
				})).Should(Equal(start + expected))
			})
		})
	})

})
