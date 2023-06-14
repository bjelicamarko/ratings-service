package router

import (
	"RatingsService/handlers"
	"RatingsService/metrics"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MapRoutesAndServe(handler *handlers.RatingsHandler) {
	log.Printf("Server is running on %s!\n", os.Getenv("SERVICE_NAME"))
	router := mux.NewRouter()

	if os.Getenv("HOST") != "localhost" && os.Getenv("MONITORING") != "OFF" { // ako nije localhost i ako nije ukljucen monitoring
		log.Println("Server is running with tracing.")
	} else {
		log.Println("Server is running without tracing.")
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
