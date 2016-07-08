package model

import "github.com/serenize/snaker"

//Table table name model
type Table struct {
	Original string
	Capital  string
}

//NewTable constructor for Table
func NewTable(tableName string) *Table {
	capital := snaker.SnakeToCamel(tableName)
	return &Table{
		Original: tableName,
		Capital:  capital,
	}
}

//NewTables multipl constructor
func NewTables(tableNames []string) []*Table {
	ret := make([]*Table, len(tableNames))
	for i, v := range tableNames {
		ret[i] = NewTable(v)
	}
	return ret
}
