package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var rateUsd = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "current_rate_coin_usd",
		Help: "Current exchange rate",
	},
	[]string{
		"coin",
	},
)

func recordMetrics() {
	coins := strings.Split(os.Getenv("COINS"), ",")
	for {
		for _, coin := range coins {
			bitcoin := fetch("rates", coin)
			if s, err := strconv.ParseFloat(bitcoin.Data.Rateusd, 64); err == nil {
				rateUsd.WithLabelValues(coin).Set(s)
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func main() {
	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
