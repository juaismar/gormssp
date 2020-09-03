package SSP

import (
	"github.com/juaismar/gormssp/test/dbs/postgres"
	"github.com/juaismar/gormssp/test/dbs/sqlite"

	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
)

type Controller struct {
	Params map[string]string
}

func (c *Controller) GetString(key string, def ...string) string {
	return c.Params[key]
}

var _ = Describe("Test SQLITE", func() {
	db := sqlite.OpenDB()

	ComplexFunctionTest(db)
	//TODO uncoment when work
	//RegExpTest(db)
	SimplexFunctionTest(db)
})

var _ = Describe("Test POSTGRES", func() {
	db := postgres.OpenDB()

	ComplexFunctionTest(db)
	RegExpTest(db)
	SimplexFunctionTest(db)
})

var _ = Describe("Test aux fuctions", func() {
	FunctionsTest()
})
