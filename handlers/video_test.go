package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideo(t *testing.T) {
	w := httptest.NewRecorder()
	id := "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345"
	req := httptest.NewRequest("GET", "/video", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": id,
	})

	handleVideoRequest(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	var data videoInfo
	if err = json.Unmarshal(jsonString, &data); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}

	assert.Equal(t, id, data.ID)
}
