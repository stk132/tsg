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

//Tpl template implements
type Tpl struct {
	tableTpl *pongo2.Template
}

//NewTpl constructor for Tpl
func NewTpl() *Tpl {
	tpl, _ := pongo2.FromString(tplString)
	return &Tpl{tableTpl: tpl}
}

//Gen generate from template
func (t *Tpl) Gen(tableNames []*model.Table) (string, error) {
	return t.tableTpl.Execute(pongo2.Context{"tables": tableNames})
}
