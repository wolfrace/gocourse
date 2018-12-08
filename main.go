package main

import (
	"net/http"
	"os"

	"github.com/echernukha/gocourse/simplevideoserver/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	serverUrl := ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	router := handlers.Router()
	log.Fatal(http.ListenAndServe(serverUrl, router))
}
