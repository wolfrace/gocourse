package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleVideoRequest(w http.ResponseWriter, _ *http.Request) {
	data := videoInfo{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
