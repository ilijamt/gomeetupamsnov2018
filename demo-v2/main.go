package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	prometheus.MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this service",
		ConstLabels: map[string]string{
			"version": "v1.2.51",
			"service": "demo-v2",
		},
	}))

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		panic(err)
	}
}
