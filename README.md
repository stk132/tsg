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
	VUsers Users
)

func init() {

	VUsers = Users{
		original: "users",

		ID: Column{name: "id"},

		Name: Column{name: "name"},
	}

}

type Column struct {
	name string
}

func (c Column) N() string {
	return c.name
}

type Users struct {
	original string

	ID Column

	Name Column
}

func (t Users) N() string {
	return t.original
}

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
dbr.Select(VUsers.Name.N()).From(VUsers.N()).Load(&m)
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
      --output-dir="."       generate file output dir
      --output-filename="const_tables.go"
                             generated filename
      --package-name="main"  generated file's pacakge name
```


## Supported RDB

postgresql only now...
