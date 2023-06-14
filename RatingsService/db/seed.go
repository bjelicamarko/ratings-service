package db

import (
	"RatingsService/models"

	"gorm.io/gorm"
)

var Ratings = []models.Rating{
	{
		Model:   gorm.Model{},
		Mark:    5,
		Comment: "onako",
	},
	{
		Model:   gorm.Model{},
		Mark:    4,
		Comment: "odlicno",
	},
}
