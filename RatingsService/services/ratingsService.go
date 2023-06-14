package services

import "RatingsService/repositories"

type RatingsService struct {
	repository *repositories.RatingsRepository
}

func NewRatingsService(repository *repositories.RatingsRepository) *RatingsService {
	return &RatingsService{repository}
}
