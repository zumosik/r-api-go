package database

import (
	"log"

	"github.com/zumosik/r-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "postgres://obojeyip:VIQYwmG0UbDurtc2gS2gMq_Mt1bMW0op@snuffleupagus.db.elephantsql.com/obojeyip"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err)
	}

	// DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Connected to db")
	log.Println("Running migrations!")
	DB.AutoMigrate(&models.User{})
}
