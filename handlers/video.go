package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleVideoRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] //Для запросов вида /video/{ID}

	if id == "" {
		http.Error(w, "invalid video id", http.StatusBadRequest)
		return
	}

	data := videoInfo{
		ID:        id,
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: fmt.Sprintf("/content/%s/screen.jpg", id),
	}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
