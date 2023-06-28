package router

import (
	"RatingsService/handlers"
	"RatingsService/metrics"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func settingTotalRequestsAndVisitors(r *http.Request) {
	metrics.TotalRequests.Inc()
	if !metrics.VisitorExists(r) {
		metrics.UniqueVisitors.Inc()
	}
}

func addingTraffics(r *http.Request, w *http.ResponseWriter) {
	if r != nil {
		metrics.TotalTrafficRequests.Add(float64(r.ContentLength))
	}
	f, _ := strconv.ParseFloat((*w).Header().Get("Content-Length"), 64)
	metrics.TotalTrafficResponses.Add(f)
}

func MapRoutesAndServe(handler *handlers.RatingsHandler) {
	log.Printf("Server is running on %s!\n", os.Getenv("SERVICE_NAME"))
	router := mux.NewRouter()

	if os.Getenv("HOST") != "localhost" && os.Getenv("MONITORING") != "OFF" { // ako nije localhost i ako nije ukljucen monitoring
		log.Println("Server is running with tracing.")

		router.HandleFunc("/ratings/add-accommodation-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/add-accommodation-rating")
			defer span.Finish()

			handler.AddAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPost)

		router.HandleFunc("/ratings/update-accommodation-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/update-accommodation-rating")
			defer span.Finish()

			handler.UpdateAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPut)

		router.HandleFunc("/ratings/accommodation-ratings/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/accommodation-ratings/{id}")
			defer span.Finish()

			handler.DeleteAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodDelete)

		router.HandleFunc("/ratings/get-ratings-for-accommodation/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/get-ratings-for-accommodation/{id}")
			defer span.Finish()

			handler.GetAccommodationsRatings(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodGet)

		// host
		router.HandleFunc("/ratings/add-host-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/add-host-rating")
			defer span.Finish()

			handler.AddHostRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPost)

		router.HandleFunc("/ratings/update-host-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/update-host-rating")
			defer span.Finish()

			handler.UpdateHostRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPut)

		router.HandleFunc("/ratings/host-ratings/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/host-ratings/{id}")
			defer span.Finish()

			handler.DeleteHostRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodDelete)

		router.HandleFunc("/ratings/get-ratings-for-host/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/get-ratings-for-host/{id}")
			defer span.Finish()

			handler.GetHostsRatings(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodGet)

	} else {
		log.Println("Server is running without tracing.")

		router.HandleFunc("/ratings/add-accommodation-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.AddAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPost)

		router.HandleFunc("/ratings/update-accommodation-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.UpdateAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPut)

		router.HandleFunc("/ratings/accommodation-ratings/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.DeleteAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodDelete)

		router.HandleFunc("/ratings/get-ratings-for-accommodation/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.GetAccommodationsRatings(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodGet)

		// host
		router.HandleFunc("/ratings/add-host-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.AddHostRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPost)

		router.HandleFunc("/ratings/update-host-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.UpdateHostRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPut)

		router.HandleFunc("/ratings/host-ratings/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.DeleteHostRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodDelete)

		router.HandleFunc("/ratings/get-ratings-for-host/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			handler.GetHostsRatings(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodGet)
	}

	router.Handle("/metrics", promhttp.Handler())

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.NotFoundRequests.Inc()
		metrics.TotalRequests.Inc()
		w.WriteHeader(http.StatusNotFound)
		log.Println("404 page not found")
	})

	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			metrics.TotalRequests.Set(0)
			metrics.SuccessfulRequests.Set(0)
			metrics.UnsuccessfulRequests.Set(0)
			metrics.NotFoundRequests.Set(0)
		}
	}()

	http.ListenAndServe(":8084", router)
}
