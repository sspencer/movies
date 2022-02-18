package main

import (
	"fmt"
	"log"
	"math/rand"
	"movies/server"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello Movies")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	s := server.NewMovieServer()

	r := httprouter.New()

	r.GET("/", homePage)
	r.GET("/movies/:id", s.GetMovie)
	r.GET("/movies", s.GetAllMovies)
	r.POST("/movies", s.CreateMovie)

	addr := ":8080"
	log.Println("starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
