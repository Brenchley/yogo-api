package config

import "database/sql"

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := ""
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
