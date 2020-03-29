package main

import (
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	log.Info("Starting up process")
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/healtz", handleHealthz)
	http.HandleFunc("/readiness", handleReadiness)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Process terminated: %v", err)
	}
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(3)

	if n != 0 {
		log.Info("Imitate process problem")
		time.Sleep(time.Second * 5)
	}

	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {

}
