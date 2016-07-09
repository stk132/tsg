package model

import "github.com/serenize/snaker"

//Column column name model
type Column struct {
	Original string
	Capital  string
}

//NewColumn constructor for Column
func NewColumn(columnName string) *Column {
	capital := snaker.SnakeToCamel(columnName)
	return &Column{
		Original: columnName,
		Capital:  capital,
	}
}

//NewColumns multiple constructor
func NewColumns(columnNames []string) []*Column {
	ret := make([]*Column, len(columnNames))
	for i, v := range columnNames {
		ret[i] = NewColumn(v)
	}
	return ret
}
