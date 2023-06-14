package repositories

import "gorm.io/gorm"

type RatingsRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RatingsRepository {
	return &RatingsRepository{db}
}
