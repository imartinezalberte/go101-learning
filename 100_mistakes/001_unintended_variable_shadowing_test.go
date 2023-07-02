package mistakes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go101-learning/100_mistakes"
)

var _ = Describe("001UnintendedVariableShadowing", func() {
	
	Describe("buildClient don't return the right client", func() {
		var tracing bool

		When("tracing is enabled", func() {
			BeforeEach(func() {
				tracing = true 
			})

			It("Should return empty structure", func() {
				Ω(mistakes.BuildClient(tracing)).Should(BeNil())
			})
		})

		When("tracing is not enabled", func() {
			BeforeEach(func() {
				tracing = false
			})

			It("Should return empty structure", func() {
				Ω(mistakes.BuildClient(tracing)).Should(BeNil())
			})
		})

		When("tracing is enabled and returns an error", func() {
			BeforeEach(func() {
				tracing = true
			})

			It("Should return error", func() {
				Ω(mistakes.BuildClientWithTracingErr(tracing)).Error().Should(MatchError(mistakes.ErrCreateClientWithTracing))
			})
		})

		When("tracing is not enabled and returns an error", func() {
			BeforeEach(func() {
				tracing = false
			})

			It("Should return error", func() {
				Ω(mistakes.BuildClientDefaultErr(tracing)).Error().Should(MatchError(mistakes.ErrCreateDefaultClient))
			})
		})
	})

	Describe("buildClient1 return the right client", func() {
		var (
			buildClient func(bool) (*mistakes.Client, error)
			tracing bool
		)

		BeforeEach(func() {
			buildClient = mistakes.BuildClientGen1(mistakes.CreateClientWithTracing, mistakes.CreateDefaultClient)
		})
		
		When("tracing is true", func() {
			BeforeEach(func() {
				tracing = true
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})

		When("tracing is false", func() {
			BeforeEach(func() {
				tracing = false
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})
	})

	Describe("buildClient2 return the right client", func() {
		var (
			buildClient func(bool) (*mistakes.Client, error)
			tracing bool
		)

		BeforeEach(func() {
			buildClient = mistakes.BuildClientGen2(mistakes.CreateClientWithTracing, mistakes.CreateDefaultClient)
		})
		
		When("tracing is true", func() {
			BeforeEach(func() {
				tracing = true
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})

		When("tracing is false", func() {
			BeforeEach(func() {
				tracing = false
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})
	})

	Describe("buildClient3 return the right client", func() {
		var (
			buildClient func(bool) (*mistakes.Client, error)
			tracing bool
		)

		BeforeEach(func() {
			buildClient = mistakes.BuildClientGen3(mistakes.CreateClientWithTracing, mistakes.CreateDefaultClient)
		})
		
		When("tracing is true", func() {
			BeforeEach(func() {
				tracing = true
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})

		When("tracing is false", func() {
			BeforeEach(func() {
				tracing = false
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})
	})


	Describe("buildClient4 return the right client", func() {
		var (
			buildClient func(bool) (*mistakes.Client, error)
			tracing bool
		)

		BeforeEach(func() {
			buildClient = mistakes.BuildClientGen4(mistakes.CreateClientWithTracing, mistakes.CreateDefaultClient)
		})
		
		When("tracing is true", func() {
			BeforeEach(func() {
				tracing = true
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})

		When("tracing is false", func() {
			BeforeEach(func() {
				tracing = false
			})

			It("buildClient return a not nil client", func() {
				Ω(buildClient(tracing)).ShouldNot(BeNil())
			})
		})
	})
})
