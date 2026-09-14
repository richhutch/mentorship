package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/richhutch/mentorship/go/http/store"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	s := store.NewInMemoryPlayerStore()
	server := NewPlayerServer(s)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
