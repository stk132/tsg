package loader

import (
	"fmt"
	//blank import for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const myTableQuery = `
SELECT
	table_name
FROM information_schema.tables
where table_schema = ?
`
const myColumnQuery = `
SELECT
	column_name
FROM information_schema.columns
WHERE table_schema = ? AND table_name = ?
ORDER BY ordinal_position
`

//MyLoader loader for MySQL
type MyLoader struct {
	sess     dbr.SessionRunner
	database string
}

//NewMyLoader constructor for MyLoader
func NewMyLoader(p *Param) (*MyLoader, error) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", p.user, p.pass, p.host, p.port, p.database)
	conn, err := dbr.Open("mysql", connectStr, nil)
	if err != nil {
		return nil, err
	}
	sess := conn.NewSession(nil)
	return &MyLoader{sess: sess, database: p.database}, nil
}

//TableNames get table name list
func (m *MyLoader) TableNames() ([]string, error) {
	var ret []string
	if _, err := m.sess.SelectBySql(myTableQuery, m.database).Load(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//ColumnNames get column name list
func (m *MyLoader) ColumnNames(tableName string) ([]string, error) {
	var ret []string
	if _, err := m.sess.SelectBySql(myColumnQuery, m.database, tableName).Load(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}
