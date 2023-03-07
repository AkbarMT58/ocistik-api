package config

import "database/sql"

func ConnectDB() *sql.DB {
	dbDriver, dbUser, dbPass, dbName := "mysql", "root", "", "ocistik"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	return db
}
