# tsg

tsg is golang source code generator for type safe query building.

generated source code is used by query buildr (ex. dbr)

## example

if exist table like below

``` sql
CREATE TABLE users (
  id serial,
  name text
)
```

tsg generate source code like below

``` go
package main

import "fmt"

var (
  //VUsers global variable for getting table, column name
  VUsers Users
)

func init() {

	VUsers = Users{
		original: "users",

		ID: Column{name: "id"},

		Name: Column{name: "name"},
	}

}

//Column struct that represents table column
type Column struct {
	name string
}

//N return column name
func (c Column) N() string {
	return c.name
}

//As return aliasName like "columnName AS aliasName"
func (c Column) As(aliasName string) string {
	return fmt.Sprintf("%s AS %s", c.N(), aliasName)
}

//Users struct that represents table "users"
type Users struct {
	original string

	ID Column

	Name Column
}

//N return table name
func (t Users) N() string {
	return t.original
}

//A return struct that has aliasName specified
func (t Users) A(aliasName string) Users {
	return Users{
		original: aliasName,

		ID: Column{name: fmt.Sprintf("%s.%s", aliasName, "id")},

		Name: Column{name: fmt.Sprintf("%s.%s", aliasName, "name")},
	}
}

```

so, using with query builder like below (using dbr)

``` go
sess.Select(VUsers.Name.N()).From(VUsers.N()).Load(&m)
```


## Usage

```
usage: tsg --user=USER --pass=PASS --database=DATABASE [<flags>]

Flags:
      --help                 Show context-sensitive help (also try --help-long
                             and --help-man).
  -u, --user=USER            database user
      --pass=PASS            password
  -h, --host="localhost"     database host
  -p, --port="5432"          database port
  -d, --database=DATABASE    database name
      --dbtype="Postgres"    database type
      --output-dir="."       generate file output dir
      --output-filename="const_tables.go"
                             generated filename
      --package-name="main"  generated file's pacakge name
```


## Supported RDB

- postgresql (dbtype: Postgres)
- mysql (dbtype: MySQL)
