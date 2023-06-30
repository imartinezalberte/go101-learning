package chapter6_test

import (
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter6"
)

var _ = Describe("Airport", func() {
	var input []chapter6.Entry

	BeforeEach(func() {
		buf, err := os.ReadFile("./airlines.json")
		Ω(err).ToNot(HaveOccurred())

		Ω(json.Unmarshal(buf, &input)).ToNot(HaveOccurred())
	})

	When("Everything goes fine", func() {
		It("", func() {
			Ω(chapter6.Task(input)).To(Equal(213848.88333333333))
		})
	})
})
