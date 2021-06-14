package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var (
	countEx = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gauge_ex_total",
			Help: "Number of counter-ex .",
		},
		[]string{"label1", "label2"},
	)
	gaugeEx = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
		Name: "gauge_ex",
		Help: "Current gauge example.",
	},
	[]string{"label1", "label2"},
	)
)


func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(countEx)
	prometheus.MustRegister(gaugeEx)
}

func main() {
	gaugeEx.With(prometheus.Labels{"label1":"hello", "label2":"wold"}).Set(54.4)
	countEx.With(prometheus.Labels{"label1":"hello", "label2":"wold"}).Inc()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}