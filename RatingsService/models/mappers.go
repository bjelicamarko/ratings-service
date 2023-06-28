package models

import "gorm.io/gorm"

func (rating *AccommodationRating) ToAccommodationRatingDTO() AccommodationRatingDTO {
	return AccommodationRatingDTO{
		Id:              rating.ID,
		Mark:            rating.Mark,
		Comment:         rating.Comment,
		GuestId:         rating.GuestId,
		HostId:          rating.HostId,
		AccommodationId: rating.AccommodationId,
	}
}

func (rating *AccommodationRating) ToAccommodationRatingForViewDTO() AccommodationRatingForViewDTO {
	accView := AccommodationRatingForViewDTO{
		Id:              rating.ID,
		Mark:            rating.Mark,
		Comment:         rating.Comment,
		GuestId:         rating.GuestId,
		HostId:          rating.HostId,
		AccommodationId: rating.AccommodationId,
	}
	if rating.UpdatedAt.IsZero() {
		accView.DateAdded = rating.UpdatedAt
	} else {
		accView.DateAdded = rating.CreatedAt
	}
	return accView
}

func (ratingDTO *AccommodationRatingDTO) ToAccommodationRating() AccommodationRating {
	return AccommodationRating{
		Model:           gorm.Model{},
		Mark:            ratingDTO.Mark,
		Comment:         ratingDTO.Comment,
		GuestId:         ratingDTO.GuestId,
		HostId:          ratingDTO.HostId,
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

func (rating *HostRating) ToHostRatingForViewDTO() HostRatingForViewDTO {
	accView := HostRatingForViewDTO{
		Id:      rating.ID,
		Mark:    rating.Mark,
		Comment: rating.Comment,
		GuestId: rating.GuestId,
		HostId:  rating.HostId,
	}
	if rating.UpdatedAt.IsZero() {
		accView.DateAdded = rating.UpdatedAt
	} else {
		accView.DateAdded = rating.CreatedAt
	}
	return accView
}
