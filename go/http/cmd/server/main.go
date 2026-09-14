package main

import (
	"log"
	"net/http"

	"github.com/richhutch/mentorship/go/http/server"
	"github.com/richhutch/mentorship/go/http/store"
)

func main() {
	s := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", s))
}
