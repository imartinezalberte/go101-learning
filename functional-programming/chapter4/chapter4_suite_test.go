package chapter4_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChapter4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chapter4 Suite")
}
