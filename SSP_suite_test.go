package ssp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSSP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SSP Suite")
}
