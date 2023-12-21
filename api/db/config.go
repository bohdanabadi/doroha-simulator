package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	port := mustGetInt(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	DB, err = gorm.Open(postgres.Open(connStr))
	//DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to obtain database handle: %v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Successfully connected to the database")
}

func mustGetInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid integer value: %s", s)
	}
	return v
}
