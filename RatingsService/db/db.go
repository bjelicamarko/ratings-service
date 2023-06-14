package db

import (
	"RatingsService/config"
	"RatingsService/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var cfg = config.ReturnConfig()
	i := 0

	for i <= 5 {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Println("Error connecting to db.")
			i++
			time.Sleep(30 * time.Second)
			continue
		} else {
			log.Println("Database connection successfully created.")
		}

		db.Migrator().DropTable("ratings")
		db.Migrator().AutoMigrate(&models.Rating{})

		for _, rating := range Ratings {
			db.Create(&rating)
		}

		return db
	}
	return nil
}
