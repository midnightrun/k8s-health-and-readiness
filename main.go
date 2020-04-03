package main

import (
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var readyToServe bool

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	log.Info("Starting up process")
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/healthz", handleHealthz)
	http.HandleFunc("/readiness", handleReadiness)
	http.HandleFunc("/toggle", handleToggle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Process terminated: %v", err)
	}
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(3)

	if n != 0 {
		log.Info("Imitate connection problem")
		time.Sleep(time.Second * 5)
	}

	w.WriteHeader(http.StatusOK)
	return
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	if readyToServe {
		log.Info("Imitate processing problem")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func handleToggle(w http.ResponseWriter, r *http.Request) {
	readyToServe = !readyToServe

	w.WriteHeader(http.StatusOK)
	return
}
