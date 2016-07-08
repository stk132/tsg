package loader

import (
	"fmt"

	"github.com/gocraft/dbr"
)

//TypePostgres loader type postgres
const TypePostgres = "Postgres"

//Loader metadata loader interface
type Loader interface {
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
