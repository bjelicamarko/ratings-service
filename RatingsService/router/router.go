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
	} else {
		log.Println("Server is running without tracing.")

		router.HandleFunc("/add-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/reservations")
			defer span.Finish()

			handler.AddAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPost)

		router.HandleFunc("/update-rating", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/update-rating")
			defer span.Finish()

			handler.UpdateAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodPut)

		router.HandleFunc("/ratings/{id}", func(w http.ResponseWriter, r *http.Request) {
			settingTotalRequestsAndVisitors(r)

			span, _ := opentracing.StartSpanFromContext(r.Context(), "/ratings/{id}")
			defer span.Finish()

			handler.DeleteAccommodationRating(w, r)

			addingTraffics(r, &w)
		}).Methods(http.MethodDelete)
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
