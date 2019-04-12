package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var db = &DB{}

func ConnectSql(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("mysql", dbSource)

	if err != nil {
		panic(err)
	}
	defer d.Close()
	db.SQL = d
	return db, nil
}
