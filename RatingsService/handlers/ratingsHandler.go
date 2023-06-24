package handlers

import (
	"RatingsService/metrics"
	"RatingsService/models"
	"RatingsService/services"
	"RatingsService/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func returnError(err error, responseStatus int, body interface{}, w *http.ResponseWriter) {
	metrics.UnsuccessfulRequests.Inc()
	log.Println(err.Error())
	(*w).WriteHeader(http.StatusBadRequest)
	json.NewEncoder(*w).Encode(body)
}

func returnResponse(logMessage string, body []byte, w *http.ResponseWriter) {
	metrics.SuccessfulRequests.Inc()
	log.Println(logMessage)

	contentLength := fmt.Sprintf("%d", len(body))
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Content-Length", contentLength)
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(body)
}

type RatingsHandler struct {
	service *services.RatingsService
}

func NewRatingsHandler(service *services.RatingsService) *RatingsHandler {
	return &RatingsHandler{service}
}

func (ah *RatingsHandler) AddAccommodationRating(w http.ResponseWriter, r *http.Request) {
	log.Println("AddAccommodationRating is called.")

	claims, err := utils.GetClaimsFromHeader(r)
	if err != nil {
		returnError(err, http.StatusBadRequest, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	var ratingDTO models.AccommodationRatingDTO
	err = json.NewDecoder(r.Body).Decode(&ratingDTO)
	if err != nil {
		returnError(err, http.StatusBadRequest, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}
	createdRating, err := ah.service.AddAccommodationRating(&ratingDTO, claims.Id)
	if err != nil {
		returnError(err, http.StatusBadRequest, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	body, err := json.Marshal(models.AccommodationRatingDTOMessage{Message: "Initiated adding of rating.", AccommodationRatingDTO: *createdRating})
	if err != nil {
		returnError(err, http.StatusInternalServerError, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	returnResponse("Initiated adding of rating", body, &w)
}

func (ah *RatingsHandler) DeleteAccommodationRating(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteAccommodationRating is called.")

	vars := mux.Vars(r)
	accommRatingId := vars["id"]

	ratingID, err := strconv.ParseUint(accommRatingId, 10, 32)
	if err != nil {
		returnError(err, http.StatusBadRequest, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	deletedRating, err := ah.service.DeleteAccommodationRating(uint(ratingID))
	if err != nil {
		returnError(err, http.StatusBadRequest, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	body, err := json.Marshal(models.AccommodationRatingDTOMessage{Message: "Deleted rating", AccommodationRatingDTO: *deletedRating})
	if err != nil {
		returnError(err, http.StatusInternalServerError, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	returnResponse("Deleted rating", body, &w)
}

func (uh *RatingsHandler) UpdateAccommodationRating(w http.ResponseWriter, r *http.Request) {
	var updateRatingDTO models.UpdateAccommodationRatingDTO
	json.NewDecoder(r.Body).Decode(&updateRatingDTO)

	log.Println("Update accommodation rating is called.")
	updatedRating, err := uh.service.UpdateAccommodationRating(&updateRatingDTO)

	if err != nil {
		returnError(err, http.StatusBadRequest, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	obj := models.AccommodationRatingDTOMessage{Message: "Update of accommodation rating succeeded.", AccommodationRatingDTO: *updatedRating}
	body, err := json.Marshal(obj)
	if err != nil {
		returnError(err, http.StatusInternalServerError, models.AccommodationRatingDTOMessage{Message: err.Error()}, &w)
		return
	}

	returnResponse("Update of accommodation rating succeeded", body, &w)
}
