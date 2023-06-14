package models

import "gorm.io/gorm"

func (rating *Rating) ToRatingDTO() RatingDTO {
	return RatingDTO{
		Id:      rating.ID,
		Mark:    rating.Mark,
		Comment: rating.Comment,
	}
}

func (ratingDTO *RatingDTO) ToRating() Rating {
	return Rating{
		Model:   gorm.Model{},
		Mark:    ratingDTO.Mark,
		Comment: ratingDTO.Comment,
	}
}
