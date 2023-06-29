package chapter4_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"github.com/imartinezalberte/go101-learning/functional-programming/chapter4"
)

var _ = Describe("Hotdog", func() {

	When("Substracting enough to stay positive in your bank account", func() {
		var (
			sender   chapter4.CreditCard[chapter4.DNI]
			receiver chapter4.CreditCard[uuid.UUID]
		)

		BeforeEach(func() {
			sender = chapter4.CreditCard[chapter4.DNI]{
				ID:     uuid.New(),
				Credit: 100,
				DNI:    "123456789",
			}
			receiver = chapter4.CreditCard[uuid.UUID]{
				ID:     uuid.New(),
				Credit: 10,
				DNI:    uuid.New(),
			}

			_, _ = sender, receiver
		})

		It("Checking that original sender is not modified and everything goes fine", func() {
			Expect(sender.Order(chapter4.HotDog{})).Should(MatchFields(IgnoreExtras, Fields{
				"Credit": Equal(sender.Credit-chapter4.HotDogPrice),
			}))
		})
	})
})
