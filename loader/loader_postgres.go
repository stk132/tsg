package loader

import (
	"fmt"

	"github.com/gocraft/dbr"
	//blank import for postgresql driver
	_ "github.com/lib/pq"
)

const pgTableQuery = `
  SELECT
	 c.relname AS table_name
  FROM pg_class c
  JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
  WHERE n.nspname = 'public' AND c.relkind = 'r'
`

const pgColumnQuery = `
SELECT
	a.attname AS column_name
FROM pg_attribute a
JOIN ONLY pg_class c ON c.oid = a.attrelid
JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid AND a.attnum = ANY(ct.conkey) AND ct.contype IN('p', 'u')
LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum
WHERE a.attisdropped = false
	AND n.nspname = 'public'
	AND c.relname = ?
	AND a.attnum > 0
ORDER BY a.attnum
`

//PgLoader Loader implements for Postgres
type PgLoader struct {
	sess dbr.SessionRunner
}

//NewPgLoader constructor for PgLoader
func NewPgLoader(p *Param) (*PgLoader, error) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", p.user, p.pass, p.host, p.port, p.database)
	conn, err := dbr.Open("postgres", connectStr, nil)
	if err != nil {
		return nil, err
	}

	sess := conn.NewSession(nil)
	return &PgLoader{sess}, nil
}

//TableNames get table name list
func (p *PgLoader) TableNames() ([]string, error) {
	var ret []string
	_, err := p.sess.SelectBySql(pgTableQuery).Load(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//ColumnNames get column name list
func (p *PgLoader) ColumnNames(tableName string) ([]string, error) {
	var ret []string
	_, err := p.sess.SelectBySql(pgColumnQuery, tableName).Load(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
