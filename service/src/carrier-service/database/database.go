package database

import (
	"carrier-service/domain/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewConnection() (*gorm.DB, error) {
	return initializeDB()
}

func initializeDB() (*gorm.DB, error) {
	connStr := "host=carrierdb port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.AutoMigrate(&model.Carrier{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
