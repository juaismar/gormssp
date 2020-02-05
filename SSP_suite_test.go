package SSP_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSSP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SSP Suite")
	junitReporter := reporters.NewJUnitReporter(os.Getenv("CI_REPORT"))
    RunSpecsWithDefaultAndCustomReporters(t, "Awesome Suite", []Reporter{junitReporter})
	