package utils

import (
	"github.com/Tucho-Will/exemplo-go/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	_"github.com/lib/pq"
)

func ConnectDb() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("db")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database : %v", err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&domain.Patient{})

	return db
}
