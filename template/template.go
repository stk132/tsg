package template

import (
	"github.com/stk132/tsg/model"
	"gopkg.in/flosch/pongo2.v3"
)

const tplString = `
package main

const (
  {% for table in tables %}
  {{table.Capital}} = "{{table.Original}}"
  {% endfor %}
)
`

const tplString2 = `
package {{package}}

import "fmt"

var (
  {% for table in tables %}
	//V{{table.Capital}} global variable for getting table, column name
  V{{table.Capital}} {{table.Capital}}
  {% endfor %}
)

func init() {
  {% for table in tables %}
  V{{table.Capital}} = {{table.Capital}}{
    original: "{{table.Original}}",
    {% for column in table.Columns %}
    {{column.Capital}}: Column{name: "{{column.Original}}"},
    {% endfor %}
  }
  {% endfor %}
}

//Column struct that represents table column
type Column struct {
  name string
}

//N return column name
func (c Column) N() string {
  return c.name
}

{% for table in tables %}
//{{table.Capital}} struct that represents table "{{table.Original}}"
type {{table.Capital}} struct {
  original string
  {% for column in table.Columns %}
  {{column.Capital}} Column
  {% endfor %}
}

//N return table name
func (t {{table.Capital}} ) N() string {
  return t.original
}

//A return struct that has aliasName specified
func (t {{table.Capital}}) A(aliasName string) {{table.Capital}} {
  return {{table.Capital}}{
    original: aliasName,
    {% for column in table.Columns %}
    {{column.Capital}}: Column{name: fmt.Sprintf("%s.%s", aliasName, "{{column.Original}}")},
    {% endfor %}
  }
}
{% endfor %}
`

//Tpl template implements
type Tpl struct {
	tableTpl *pongo2.Template
}

//NewTpl constructor for Tpl
func NewTpl() *Tpl {
	tpl, _ := pongo2.FromString(tplString2)
	return &Tpl{tableTpl: tpl}
}

//Gen generate from template
func (t *Tpl) Gen(tableNames []*model.Table, packageName string) (string, error) {
	return t.tableTpl.Execute(pongo2.Context{"tables": tableNames, "package": packageName})
}
