package models

type AccommodationRatingDTO struct {
	Id              uint   `json:"Id"`
	Mark            int    `json:"Mark"`
	Comment         string `json:"Comment"`
	GuestId         uint   `json:"GuestId"`
	AccommodationId uint   `json:"AccommodationId"`
}

type AccommodationRatingDTOMessage struct {
	AccommodationRatingDTO AccommodationRatingDTO `json:"ratingDTO"`
	Message                string                 `json:"message"`
}

type UpdateAccommodationRatingDTO struct {
	Id      uint   `json:"Id"`
	Mark    int    `json:"Mark" validate:"min=1,max=5"`
	Comment string `json:"Comment"`
}

type HostRatingDTO struct {
	Id      uint   `json:"Id"`
	Mark    int    `json:"Mark"`
	Comment string `json:"Comment"`
	GuestId uint   `json:"GuestId"`
	HostId  uint   `json:"AccommodationId"`
}

type HostRatingDTOMessage struct {
	HostRatingDTO HostRatingDTO `json:"ratingDTO"`
	Message       string        `json:"message"`
}

type UpdateHostRatingDTO struct {
	Id      uint   `json:"Id"`
	Mark    int    `json:"Mark" validate:"min=1,max=5"`
	Comment string `json:"Comment"`
}
