package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

var readyToServe bool

type Application struct {
	Status bool
}

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	log.Info("Starting up process")
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", handleStatus)
	http.HandleFunc("/healthz", handleHealthz)
	http.HandleFunc("/readiness", handleReadiness)
	http.HandleFunc("/toggle", handleToggle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Process terminated: %v", err)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Status ...")
	fp := path.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, Application{Status: readyToServe}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Liveness ...")
	n := rand.Intn(3)

	if n != 0 {
		log.Info("Imitate connection problem")
		time.Sleep(time.Second * 5)
	}

	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Readiness ...")
	if readyToServe {
		log.Info("Imitate processing problem")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleToggle(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Toggle ...")
	readyToServe = !readyToServe

	w.WriteHeader(http.StatusOK)
}
