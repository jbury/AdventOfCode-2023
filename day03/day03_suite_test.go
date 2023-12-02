package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDay03(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day03 Suite")
}
