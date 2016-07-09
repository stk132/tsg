package loader

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/stk132/tsg/model"
)

//TypePostgres loader type postgres
const TypePostgres = "Postgres"

//Loader metadata loader interface
type Loader interface {
	Load(dbr.SessionRunner) ([]*model.Table, error)
	TableNames(dbr.SessionRunner) ([]string, error)
}

//NewLoader loader constructer
func NewLoader(loaderType string) (Loader, error) {
	switch loaderType {
	case TypePostgres:
		return &PgLoader{}, nil
	}
	return nil, fmt.Errorf("loaderType: %s is not found", loaderType)
}
