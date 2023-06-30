package chapter7_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter7"
)

var _ = Describe("X", func() {

	Describe("Playing with trees", func() {
		var input *chapter7.Node

		When("Node Sum", func() {
			BeforeEach(func() {
				//			1
				//		/   \
				//	 2		 5
				//  / \
				// 3   4
				input = &chapter7.Node{
					Value: 1,
					Left: &chapter7.Node{
						Value: 2,
						Left: &chapter7.Node{
							Value: 3,
						},
						Right: &chapter7.Node{
							Value: 4,
						},
					},
					Right: &chapter7.Node{
						Value: 5,
					},
				}
			})

			It("is iterative", func() {
				Ω(chapter7.NodeSumIt(input)).Should(Equal(15))
			})

			It("is recursive", func() {
				Ω(chapter7.NodeSumRecursive(input)).Should(Equal(15))
			})
		})

		When("NodeMax", func() {
			BeforeEach(func() {
				//			1
				//		/   \
				//	 2		 5
				//  / \
				// 3   4
				input = &chapter7.Node{
					Value: 1,
					Left: &chapter7.Node{
						Value: 2,
						Left: &chapter7.Node{
							Value: 3,
						},
						Right: &chapter7.Node{
							Value: 4,
						},
					},
					Right: &chapter7.Node{
						Value: 5,
					},
				}
			})

			It("max number is 5", func() {
				Ω(chapter7.NodeMax(input)).Should(Equal(5))
			})
		})
	})
})
