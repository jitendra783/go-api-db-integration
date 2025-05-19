package main

import (
	"flag"
	"fmt"
	"go-api/api"
	"go-api/pkgs/config"
	"go-api/pkgs/db"

	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// A counter to track the total number of HTTP requests
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status"}, // Labels: method and status
	)

	// A histogram to track the duration of HTTP requests
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Histogram of HTTP request durations",
			Buckets: prometheus.DefBuckets, // Default bucket sizes
		},
		[]string{"method"}, // Label: method (GET, POST, etc.)
	)
)

// var db *gorm.DB
func init() {
	// Register metrics with Prometheus
	println("prometheus metrics registered")
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(httpDuration)
}

func init() {
	fmt.Println("init to Oracle DB---------------------")
	db.ConnectOracleDb("")
	fmt.Println("init to Oracle DB---------------------")
}

func main() {
	log.Println("Starting main function")
	flag.Usage = func() {
		log.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	envirnment := "local"
	host := os.Getenv("SERVER_HOST")
	if host != "" {
		envirnment = "server"
	}
	config.Load(envirnment)
	if err := api.Start(); err != nil {
		log.Fatal("Failed to start server, err:", err)
	}

}
