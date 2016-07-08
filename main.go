package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	_ "github.com/serenize/snaker"
	"github.com/stk132/tsg/loader"
	"github.com/stk132/tsg/model"
	"github.com/stk132/tsg/template"
)

func main() {
	conn, err := dbr.Open("postgres", "postgres://stk132:postgres@localhost:5432/mydb?sslmode=disable", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	sess := conn.NewSession(nil)
	l, err := loader.NewLoader(loader.TypePostgres)
	if err != nil {
		fmt.Println(err)
		return
	}

	tableNames, err := l.TableNames(sess)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpl := template.NewTpl()
	tables := model.NewTables(tableNames)
	out, err := tpl.Gen(tables)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = ioutil.WriteFile("const_tables.go", []byte(out), 0660); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
	fmt.Println(tableNames)

}
