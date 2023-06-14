package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Mark    int    `gorm:"not null"`
	Comment string `gorm:"not null"`
}
