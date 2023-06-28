package chapter2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter2"
)

var _ = Describe("Todo", func() {
	Describe("Permissions", func() {
		var (
			username string
			auth     chapter2.AuthFn
		)

		BeforeEach(func() {
			// Here we can observe the power of mocking using FP. We could probably use an interface, It just depends on the case.
			auth = func(username string) bool { return username == "admin" }
		})

		When("username is allowed", func() {
			BeforeEach(func() {
				username = "notallowed"
			})

			It("username is not allowed", func() {
				Expect(chapter2.NewToDo(&chapter2.Db{auth}).Write(username, "dummy thing")).Error().Should(HaveOccurred())
			})
		})

		When("username is allowed and have writting permissions", func() {
			BeforeEach(func() {
				username = "admin"
			})

			It("username is allowed", func() {
				Expect(chapter2.NewToDo(&chapter2.Db{auth}).Write(username, "dummy thing")).Error().ShouldNot(HaveOccurred())
			})
		})
	})
})
