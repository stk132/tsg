package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/stk132/tsg/loader"
	"github.com/stk132/tsg/template"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	user     = kingpin.Flag("user", "database user").Short('u').Required().String()
	pass     = kingpin.Flag("pass", "password").Required().String()
	host     = kingpin.Flag("host", "database host").Default("localhost").Short('h').String()
	port     = kingpin.Flag("port", "database port").Default("5432").Short('p').String()
	database = kingpin.Flag("database", "database name").Short('d').Required().String()

	dir         = kingpin.Flag("output-dir", "generate file output dir").Default(".").String()
	output      = kingpin.Flag("output-filename", "generated filename").Default("const_tables.go").String()
	packageName = kingpin.Flag("package-name", "generated file's pacakge name").Default("main").String()
)

func errHandle(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	kingpin.Parse()
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *user, *pass, *host, *port, *database)
	conn, err := dbr.Open("postgres", connectStr, nil)
	if err != nil {
		errHandle(err)
	}

	sess := conn.NewSession(nil)
	l, err := loader.NewLoader(loader.TypePostgres)
	if err != nil {
		errHandle(err)
	}

	tables, err := l.Load(sess)
	if err != nil {
		errHandle(err)
	}

	tpl := template.NewTpl()
	out, err := tpl.Gen(tables, *packageName)
	if err != nil {
		errHandle(err)
	}

	if err = os.MkdirAll(*dir, os.ModeDir|0755); err != nil {
		errHandle(err)
	}

	fileName := fmt.Sprintf("%s/%s", *dir, *output)
	if err = ioutil.WriteFile(fileName, []byte(out), 0660); err != nil {
		errHandle(err)
	}

}
