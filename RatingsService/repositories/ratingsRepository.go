package repositories

import (
	"RatingsService/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RatingsRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RatingsRepository {
	return &RatingsRepository{db}
}

func (repo *RatingsRepository) CreateAccommodationRating(ratingDTO *models.AccommodationRatingDTO) (*models.AccommodationRating, error) {
	rating := ratingDTO.ToAccommodationRating()
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Model(&rating).Create(&rating)
		if result.Error != nil {
			return result.Error
		}

		message, err := createMessage("ADD_ACCOMMODATION_RATING_INITIATED", models.AddAccommodationRatingInitiated{GuestId: rating.GuestId,
			AccommodationId: rating.AccommodationId, RatingId: rating.ID})
		if err != nil {
			return err
		}

		result = repo.db.Table("messages").Create(&message)
		if result.Error != nil {
			return errors.New("error while creating event for deletion initiated")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (repo *RatingsRepository) HasUserAlreadyRatedAccommodation(guestId uint, accommodationId uint) (bool, error) {
	var resCount int64

	result := repo.db.Table("accommodation_ratings").
		Where("deleted_at IS NULL").
		Where("guest_id = ?", guestId).
		Where("accommodation_id = ?", accommodationId).
		Where("status != ?", models.REJECTED).
		Count(&resCount)

	if result.Error != nil {
		return false, result.Error
	}

	return resCount > 0, nil
}

func (repo *RatingsRepository) GetAccommodationRatingById(ratingId uint) (*models.AccommodationRating, error) {
	var rating *models.AccommodationRating

	result := repo.db.First(&rating, ratingId).Where("deleted_at IS NULL")

	if result.Error != nil {
		return nil, result.Error
	}

	if rating == nil {
		errorMessage := fmt.Sprintf("rating not found by id=%d", ratingId)
		return nil, errors.New(errorMessage)
	}

	return rating, nil
}

func (repo *RatingsRepository) DeleteAccommodationRating(id uint) (*models.AccommodationRating, error) {
	var deletedAccommodationRating models.AccommodationRating
	result := repo.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&deletedAccommodationRating)

	if result.Error != nil {
		return nil, result.Error
	}

	return &deletedAccommodationRating, nil
}

func (repo *RatingsRepository) UpdateAccommodationRating(rating *models.AccommodationRating) (*models.AccommodationRating, error) {
	result := repo.db.Model(&rating).Updates(&rating)

	return rating, result.Error
}

func (repo *RatingsRepository) GetRatingsForAccommodation(accommodationId uint) ([]*models.AccommodationRating, error) {
	var ratings []*models.AccommodationRating

	result := repo.db.Table("accommodation_ratings").
		Where("deleted_at IS NULL").
		Where("accommodation_id = ?", accommodationId).
		Where("status = ?", models.ACCEPTED).
		Find(&ratings)

	if result.Error != nil {
		return nil, result.Error
	}

	return ratings, nil
}

//// host ratings

func (repo *RatingsRepository) CreateHostRating(ratingDTO *models.HostRatingDTO) (*models.HostRating, error) {
	rating := ratingDTO.ToHostRating()
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := repo.db.Model(&rating).Create(&rating)
		if result.Error != nil {
			return result.Error
		}

		message, err := createMessage("ADD_HOST_RATING_INITIATED", models.AddHostRatingInitiated{GuestId: rating.GuestId,
			HostId: rating.HostId, RatingId: rating.ID})
		if err != nil {
			return err
		}

		result = repo.db.Table("messages").Create(&message)
		if result.Error != nil {
			return errors.New("error while creating event for create of host rating")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (repo *RatingsRepository) HasUserAlreadyRatedHost(guestId uint, hostId uint) (bool, error) {
	var resCount int64

	result := repo.db.Table("host_ratings").
		Where("deleted_at IS NULL").
		Where("guest_id = ?", guestId).
		Where("host_id = ?", hostId).
		Where("status != ?", models.REJECTED).
		Count(&resCount)

	if result.Error != nil {
		return false, result.Error
	}

	return resCount > 0, nil
}

func (repo *RatingsRepository) GetHostRatingById(ratingId uint) (*models.HostRating, error) {
	var rating *models.HostRating

	result := repo.db.First(&rating, ratingId).Where("deleted_at IS NULL")

	if result.Error != nil {
		return nil, result.Error
	}

	if rating == nil {
		errorMessage := fmt.Sprintf("rating not found by id=%d", ratingId)
		return nil, errors.New(errorMessage)
	}

	return rating, nil
}

func (repo *RatingsRepository) DeleteHostRating(id uint) (*models.HostRating, error) {
	var deletedHostRating models.HostRating
	result := repo.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&deletedHostRating)

	if result.Error != nil {
		return nil, result.Error
	}

	return &deletedHostRating, nil
}

func (repo *RatingsRepository) UpdateHostRating(rating *models.HostRating) (*models.HostRating, error) {
	result := repo.db.Model(&rating).Updates(&rating)

	return rating, result.Error
}

func (repo *RatingsRepository) GetRatingsForHost(hostId uint) ([]*models.HostRating, error) {
	var ratings []*models.HostRating

	result := repo.db.Table("host_ratings").
		Where("deleted_at IS NULL").
		Where("host_id = ?", hostId).
		Where("status = ?", models.ACCEPTED).
		Find(&ratings)

	if result.Error != nil {
		return nil, result.Error
	}

	return ratings, nil
}

////

func createMessage(messageType models.MessageType, body interface{}) (*models.Message, error) {
	marshalled, err := json.Marshal(body)

	if err != nil {
		log.Printf("error occurred while marshalling message body of type [%s]", messageType)
		return nil, err
	}

	message := models.Message{
		Model: gorm.Model{},
		Type:  messageType,
		Body:  datatypes.JSON([]byte(marshalled)),
	}

	return &message, nil
}

func (repo *RatingsRepository) PublishMessage(message *models.Message) error {
	result := repo.db.Table("messages").Create(&message)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
