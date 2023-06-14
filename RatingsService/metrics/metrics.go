package metrics

import (
	"RatingsService/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var TotalRequests = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "total_requests",
	Help: "The total number of processed requests",
})

var SuccessfulRequests = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "successful_requests",
	Help: "The total number of successful requests",
})

var UnsuccessfulRequests = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "unsuccessful_requests",
	Help: "The total number of unsuccessful requests",
})

var NotFoundRequests = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "not_found_requests",
	Help: "The total number of 404 requests",
})

var TotalTrafficRequests = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "total_traffic_requests",
	Help: "Total traffic of requests flow in GB",
})

var TotalTrafficResponses = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "total_traffic_responses",
	Help: "Total traffic of responses flow in GB",
})

var UniqueVisitorsList []string

var UniqueVisitors = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "unique_visitors",
	Help: "Number of unique visitors",
})

func VisitorExists(r *http.Request) bool {
	ip := r.Header.Get("X-Real-IP")
	timestamp := time.Now().Unix()
	userAgent := r.UserAgent()
	visitorID := fmt.Sprintf("%s-%d-%s", ip, timestamp, userAgent)

	if utils.Contains(UniqueVisitorsList, visitorID) {
		return true
	}
	UniqueVisitorsList = append(UniqueVisitorsList, visitorID)
	return false
}
