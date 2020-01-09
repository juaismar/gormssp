# gormssp

Documentation in progres on 09-02-2020

Example of use whit beego:

  import ("github.com/juaismar/gormssp")

  func (c *User) Pagination() {

    columns := make(map[int]SSP.Data)
    columns[0] = SSP.Data{Db: "name", Dt: 0, Formatter: nil}
    columns[1] = SSP.Data{Db: "role", Dt: 1, Formatter: nil}
    columns[2] = SSP.Data{Db: "email", Dt: 2, Formatter: nil}

    c.Data["json"] = SSP.Simple(c, model.ORM, "users", columns)
    c.ServeJSON()
  }

  -In SSP.Data:-
  Db: The name of field in DB
  Dt: Position in the client table.
  Formatter: Function to format the returned data


Client side and original PHP code
<br>
https://datatables.net/examples/data_sources/server_side