package SSP

import (
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Controller struct {
	Params map[string]string
}

func (c *Controller) GetString(key string, def ...string) string {
	return c.Params[key]
}

var _ = Describe("Test for SSP", func() {

	Describe("flated", func() {
		It("returns Empty", func() {

			var whereArray []string

			result := flated(whereArray)

			Expect(result).To(Equal(""))
		})
		It("returns one query", func() {

			var whereArray []string
			whereArray = append(whereArray, "number = 1")

			result := flated(whereArray)

			Expect(result).To(Equal("number = 1"))
		})
		It("returns two query", func() {

			var whereArray []string
			whereArray = append(whereArray, "number = 1")
			whereArray = append(whereArray, "name = 'John'")

			result := flated(whereArray)

			Expect(result).To(Equal("number = 1 AND name = 'John'"))
		})
	})
	Describe("search", func() {
		It("returns -1", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "")

			Expect(result).To(Equal(-1))
		})
		It("returns -1", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "instrument")

			Expect(result).To(Equal(-1))
		})
		It("returns 1", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "role")

			Expect(result).To(Equal(1))
		})
		It("returns 0", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "0")

			Expect(result).To(Equal(0))
		})

	})
})
