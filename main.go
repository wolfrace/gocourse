package main

import (
	"context"
	"github.com/echernukha/gocourse/simplevideoserver/tools"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/echernukha/gocourse/simplevideoserver/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer tools.CloseQuietly(file)
	}

	killSignalCh := getKillSignalCh()
	serverUrl := ":8000"
	srv := startServer(serverUrl)
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")

	waitForKillSignal(killSignalCh)
	_ = srv.Shutdown(context.Background())
}

func startServer(serverUrl string) *http.Server {
	router := handlers.Router()
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

func getKillSignalCh() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
