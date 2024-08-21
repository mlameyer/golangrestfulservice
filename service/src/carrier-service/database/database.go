package database

import (
	"database/sql"
	"log"
)

func NewConnection() (*sql.DB, error) {
	return getPostgreSqlConn()
}

func getPostgreSqlConn() (*sql.DB, error) {
	connStr := "user=carrierservice dbname=carrierdb sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
