package models

import "gorm.io/gorm"

func (rating *AccommodationRating) ToAccommodationRatingDTO() AccommodationRatingDTO {
	return AccommodationRatingDTO{
		Id:              rating.ID,
		Mark:            rating.Mark,
		Comment:         rating.Comment,
		GuestId:         rating.GuestId,
		AccommodationId: rating.AccommodationId,
	}
}

func (ratingDTO *AccommodationRatingDTO) ToAccommodationRating() AccommodationRating {
	return AccommodationRating{
		Model:           gorm.Model{},
		Mark:            ratingDTO.Mark,
		Comment:         ratingDTO.Comment,
		GuestId:         ratingDTO.GuestId,
		AccommodationId: ratingDTO.AccommodationId,
	}
}

func (rating *HostRating) ToHostRatingDTO() HostRatingDTO {
	return HostRatingDTO{
		Id:      rating.ID,
		Mark:    rating.Mark,
		Comment: rating.Comment,
		GuestId: rating.GuestId,
		HostId:  rating.HostId,
	}
}

func (ratingDTO *HostRatingDTO) ToHostRating() HostRating {
	return HostRating{
		Model:   gorm.Model{},
		Mark:    ratingDTO.Mark,
		Comment: ratingDTO.Comment,
		GuestId: ratingDTO.GuestId,
		HostId:  ratingDTO.HostId,
	}
}
