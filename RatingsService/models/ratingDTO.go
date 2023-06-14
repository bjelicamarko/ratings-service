package models

type RatingDTO struct {
	Id      uint   `json:"Id"`
	Mark    int    `json:"Mark"`
	Comment string `json:"Comment"`
}
