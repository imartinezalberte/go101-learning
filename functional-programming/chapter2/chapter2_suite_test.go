package chapter2_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChapter2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chapter2 Suite")
}
