package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// connection *gorm.DB
	err error
)

func SetupConnection() *gorm.DB {
	HOST := "localhost"
	USER := "igoy"
	PASS := 12345678
	PORT := 5432
	DBNAME := "bwastartup"
	dsn := fmt.Sprintf("host=%s user=%s password=%d dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", HOST, USER, PASS, DBNAME, PORT)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect db error: %s", err)
	}

	return connection
}

// func DBConnection() *gorm.DB {
// 	return connection
// }
