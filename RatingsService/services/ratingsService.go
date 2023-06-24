package services

import (
	"RatingsService/models"
	"RatingsService/repositories"
	"errors"

	"github.com/go-playground/validator/v10"
)

type RatingsService struct {
	repository *repositories.RatingsRepository
}

func NewRatingsService(repository *repositories.RatingsRepository) *RatingsService {
	return &RatingsService{repository}
}

func (rs *RatingsService) AddAccommodationRating(ratingDTO *models.AccommodationRatingDTO, guestId uint) (*models.AccommodationRatingDTO, error) {
	validate := validator.New()
	err := validate.Struct(ratingDTO)
	if err != nil {
		return nil, err
	}
	ratingDTO.GuestId = guestId

	already_rated, err := rs.repository.HasUserAlreadyRatedAccommodation(guestId, ratingDTO.AccommodationId)
	if err != nil {
		return nil, err
	}

	if already_rated {
		return nil, errors.New("user has already rated the accommodation")
	}

	rating, err := rs.repository.CreateAccommodationRating(ratingDTO)

	if err != nil {
		return nil, err
	}

	retValue := rating.ToAccommodationRatingDTO()
	return &retValue, nil
}

func (rs *RatingsService) DeleteAccommodationRating(ratingId uint) (*models.AccommodationRatingDTO, error) {
	rating, err := rs.repository.DeleteAccommodationRating(ratingId)

	if err != nil {
		return nil, err
	}

	retValue := rating.ToAccommodationRatingDTO()
	return &retValue, nil
}

func (rs *RatingsService) UpdateAccommodationRating(updateRating *models.UpdateAccommodationRatingDTO) (*models.AccommodationRatingDTO, error) {
	rating, err := rs.repository.FindAccommodationRatingById(updateRating.Id)

	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(updateRating)
	if err != nil {
		return nil, err
	}

	rating.Mark = updateRating.Mark
	rating.Comment = updateRating.Comment

	rating, err = rs.repository.UpdateAccommodationRating(rating)
	if err != nil {
		return nil, err
	}

	retValue := rating.ToAccommodationRatingDTO()
	return &retValue, nil
}
