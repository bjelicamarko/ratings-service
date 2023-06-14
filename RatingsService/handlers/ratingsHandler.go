package handlers

import "RatingsService/services"

type RatingsHandler struct {
	service *services.RatingsService
}

func NewRatingsHandler(service *services.RatingsService) *RatingsHandler {
	return &RatingsHandler{service}
}
