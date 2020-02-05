package SSP

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("flated", func() {
	Describe("NewResetToken", func() {
		It("returns Empty", func() {

			var whereArray []string

			result := flated(whereArray)

			Expect(result).To(Equal(""))
		})
	})
})