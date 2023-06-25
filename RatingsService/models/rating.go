package models

import "gorm.io/gorm"

type AccommodationRating struct {
	gorm.Model
	Mark            int        `gorm:"not null" validate:"min=1,max=5"`
	Comment         string     `gorm:"not null"`
	GuestId         uint       `gorm:"not null"`
	AccommodationId uint       `gorm:"not null"`
	Status          SagaStatus `gorm:"not null;default:'PENDING'"`
}

type HostRating struct {
	gorm.Model
	Mark    int        `gorm:"not null" validate:"min=1,max=5"`
	Comment string     `gorm:"not null"`
	GuestId uint       `gorm:"not null"`
	HostId  uint       `gorm:"not null"`
	Status  SagaStatus `gorm:"not null;default:'PENDING'"`
}
