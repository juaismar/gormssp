# gormssp

_Using Datatables pagination with golang_

### Pre-requisites 📋

* Oviously use it in a golang project
* Gorm package (https://gorm.io/)(github.com/jinzhu/gorm)
* Beego package (https://beego.me/)(github.com/astaxie/beego)

### Installation 🔧

_Install whit the next command:_

```
go get github.com/juaismar/gormssp
```

_and import the package with:_

```
import ("github.com/juaismar/gormssp")
```
## Working example 🚀

-// TODO, show a better example
```
import ("github.com/juaismar/gormssp")

func (c *User) Pagination() {

  // Array of database columns which should be read and sent back to DataTables.
  // The `db` parameter represents the column name in the database, while the `dt`
  // parameter represents the DataTables column identifier. In this case simple
  // indexes
  // Formatter is a function to customize the value of field , can be nil.
  columns := make(map[int]SSP.Data)
  columns[0] = SSP.Data{Db: "name", Dt: 0, Formatter: nil}
  columns[1] = SSP.Data{Db: "role", Dt: 1, Formatter: nil}
  columns[2] = SSP.Data{Db: "email", Dt: 2, Formatter: nil}

  // Send the data to the client
  c.Data["json"] = SSP.Simple(c, model.ORM, "users", columns)
  c.ServeJSON()
}
```

-This project is based in the PHP version of datatables pagination in https://datatables.net/examples/data_sources/server_side -

## Author ✒️

* **Juan Iscar** - (https://github.com/juaismar)

## Thanks 🎁
* All my friends at work.
* Sergio and Mario (https://github.com/mapreal19) who taught me how to program golang and showed me the wonderful world of good practices.
* Juan, Juan and Joaquin.


_Readme.md based in https://gist.github.com/Villanuevand/6386899f70346d4580c723232524d35a_

## Notes

The search dont work prperly before Nov 08 2019, this update solve the problem
https://go-review.googlesource.com/c/go/+/205897/