package mysql

import "database/sql"

func Connect(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	return
}
