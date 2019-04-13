package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var DataSource = &DB{}

func ConnectSql(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s",
		pass,
		host,
		port,
		dbname,
	)
	d, _ := sql.Open("mysql", dbSource)
	err := d.Ping()
	DataSource.SQL = d
	return DataSource, err
}
