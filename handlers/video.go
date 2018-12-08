package handlers

import (
	"fmt"
	"net/http"
)

func handleVideoRequest(w http.ResponseWriter, _ *http.Request) {
	data := `{
	"id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
	"name": "Black Retrospetive Woman",
	"duration": 15,
	"thumbnail":"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	"url":"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4"
	}`
	fmt.Fprint(w, data)
}
