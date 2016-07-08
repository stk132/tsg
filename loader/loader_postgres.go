package loader

import "github.com/gocraft/dbr"

const pgTableQuery = `
  SELECT
	 c.relname AS table_name
  FROM pg_class c
  JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
  WHERE n.nspname = 'public' AND c.relkind = 'r'
`

//PgLoader Loader implements for Postgres
type PgLoader struct {
}

//TableNames get table name list
func (p *PgLoader) TableNames(sess dbr.SessionRunner) ([]string, error) {
	var ret []string
	_, err := sess.SelectBySql(pgTableQuery).Load(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
