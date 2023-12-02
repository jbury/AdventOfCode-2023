package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDay02(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day02 Suite")
}
