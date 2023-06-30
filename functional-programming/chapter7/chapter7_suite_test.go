package chapter7_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChapter7(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chapter7 Suite")
}
