package database

import (
	"carrier-service/domain/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewConnection() (*gorm.DB, error) {
	return initializeDB()
}

//func getPostgreSqlConn() (*sql.DB, error) {
//	connStr := "postgresql://postgres:mysecretpassword@localhost:2022/carrierdb?sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//		return nil, err
//	}
//	if err = db.Ping(); err != nil {
//		return nil, err
//	}
//	CreateCarrierTable(db)
//	return db, nil
//}

func initializeDB() (*gorm.DB, error) {
	connStr := "postgresql://postgres:mysecretpassword@localhost:2022?sslmode=disable"
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

	db.Create(&model.Carrier{
		Name:      "test carrier",
		InService: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return db, nil
}

//func CreateCarrierTable(db *sql.DB) {
//	db.Query(`CREATE TABLE IF NOT EXISTS carriers (
//    id SERIAL PRIMARY KEY,
//    name VARCHAR(255) NOT NULL,
//    in_service BIT NOT NULL,
//    created_at TIMESTAMP NOT NULL,
//    updated_at TIMESTAMP NOT NULL,
//)`)
//}
