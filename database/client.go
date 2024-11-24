package database

import (
	"log"

	"github.com/JAANVI999/JWT_Auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	log.Println("Connecting to database...")
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to db")
	}
	log.Println("Connected to database")
}
func Migrate() {
	log.Println("Running database migration...")
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrated")
}
