package SSP

import (
	"encoding/hex"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("flated", func() {
	Describe("NewResetToken", func() {
		It("returns Empty", func() {

			whereArray []string

			result := SSP.flated(whereArray)

			Expect(result).To(Equal(""))
		})
	})
})