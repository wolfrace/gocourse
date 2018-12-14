package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleListRequest(w http.ResponseWriter, _ *http.Request) {
	var data []videoInfo
	data = append(data, videoInfo{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	})

	data = append(data, videoInfo{
		ID:        "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
		Name:      "Танцор",
		Duration:  92,
		Thumbnail: "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
	})

	data = append(data, videoInfo{
		ID:        "sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
		Name:      "Go Rally TEASER-HD",
		Duration:  41,
		Thumbnail: "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
	})

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
