package loader

import (
	"fmt"

	"github.com/stk132/tsg/model"
)

//TypePostgres loader type postgres
const (
	TypePostgres = "Postgres"
	TypeMySQL    = "MySQL"
)

//Loader metadata loader interface
type Loader interface {
	TableNames() ([]string, error)
	ColumnNames(string) ([]string, error)
}

//Param parameter struct for loader construction
type Param struct {
	user     string
	pass     string
	host     string
	port     string
	database string
}

//NewParam constructor for Param
func NewParam(user, pass, host, port, database string) *Param {
	return &Param{
		user,
		pass,
		host,
		port,
		database,
	}
}

//NewLoader loader constructer
func NewLoader(loaderType string, p *Param) (Loader, error) {
	var l Loader
	var err error

	switch loaderType {
	case TypePostgres:
		l, err = NewPgLoader(p)
	case TypeMySQL:
		l, err = NewMyLoader(p)
	default:
		return nil, fmt.Errorf("loaderType: %s is not found", loaderType)
	}
	return l, err
}

//Load load schema data
func Load(l Loader) ([]*model.Table, error) {
	tableNames, err := l.TableNames()
	if err != nil {
		return nil, err
	}

	elements := make([]*model.Elem, len(tableNames))
	for i, v := range tableNames {
		columnNames, err := l.ColumnNames(v)
		if err != nil {
			return nil, err
		}
		elements[i] = &model.Elem{
			TableName:   v,
			ColumnNames: columnNames,
		}
	}
	return model.NewTables(elements), nil
}
