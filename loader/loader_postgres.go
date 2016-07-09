package loader

import (
	"github.com/gocraft/dbr"
	"github.com/stk132/tsg/model"
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
}

//Load load schema data
func (p *PgLoader) Load(sess dbr.SessionRunner) ([]*model.Table, error) {
	tableNames, err := p.TableNames(sess)
	if err != nil {
		return nil, err
	}

	elements := make([]*model.Elem, len(tableNames))
	for i, v := range tableNames {
		columnNames, err := p.ColumnNames(sess, v)
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

//TableNames get table name list
func (p *PgLoader) TableNames(sess dbr.SessionRunner) ([]string, error) {
	var ret []string
	_, err := sess.SelectBySql(pgTableQuery).Load(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//ColumnNames get column name list
func (p *PgLoader) ColumnNames(sess dbr.SessionRunner, tableName string) ([]string, error) {
	var ret []string
	_, err := sess.SelectBySql(pgColumnQuery, tableName).Load(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
