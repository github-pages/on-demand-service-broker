package boshlinks_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoshlinks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Boshlinks Suite")
}
