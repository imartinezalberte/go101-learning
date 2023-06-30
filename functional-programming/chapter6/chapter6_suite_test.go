package chapter6_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChapter6(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chapter6 Suite")
}
