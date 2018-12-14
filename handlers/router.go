package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/list", handleListRequest).Methods(http.MethodGet)
	s.HandleFunc("/video/{id}", handleVideoRequest).Methods(http.MethodGet)

	return logMiddleware(r)
}
