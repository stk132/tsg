package model

import "github.com/serenize/snaker"

//Elem hogehogehoge
type Elem struct {
	TableName   string
	ColumnNames []string
}

//Table table name model
type Table struct {
	Original string
	Capital  string
	Columns  []*Column
}

//NewTable constructor for Table
func NewTable(tableName string, columnNames []string) *Table {
	capital := snaker.SnakeToCamel(tableName)
	columns := NewColumns(columnNames)
	return &Table{
		Original: tableName,
		Capital:  capital,
		Columns:  columns,
	}
}

//NewTables multipl constructor
func NewTables(elements []*Elem) []*Table {
	ret := make([]*Table, len(elements))
	for i, v := range elements {
		ret[i] = NewTable(v.TableName, v.ColumnNames)
	}
	return ret
}
