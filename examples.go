/* func (c *Userloged) Pagination() {

	var ssp SSP.SSP

	columns := make(map[int]SSP.Data)
	columns[0] = SSP.Data{Db: "name", Dt: 0, Formatter: nil}
	columns[1] = SSP.Data{Db: "role", Dt: 1, Formatter: nil}
	columns[2] = SSP.Data{Db: "email", Dt: 2, Formatter: nil}

	c.Data["json"] = ssp.Simple(c, model.ORM, "users", columns)
	c.ServeJSON()
}*/