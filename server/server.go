package server

import (
	"encoding/json"
	"movies/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MovieServer struct {
	movies *services.MovieService
}

func NewMovieServer() *MovieServer {
	return &MovieServer{movies: services.NewMovieService()}
}

func (s *MovieServer) GetMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		replyError(w, r, http.StatusBadRequest, "id must be a number")
		return
	}
	movie, err := s.movies.GetMovie(id)
	if err != nil {
		replyError(w, r, http.StatusNotFound, err.Error())
		return
	}

	replyJSON(w, r, http.StatusOK, movie)
}

func (s *MovieServer) GetAllMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies := s.movies.GetMovies()
	replyJSON(w, r, http.StatusOK, movies)
}

func (s *MovieServer) CreateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var incoming struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
	}

	r.Body = http.MaxBytesReader(w, r.Body, 102_400)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&incoming)

	if err != nil {
		replyError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	movie, err := s.movies.AddMovie(incoming.Title, incoming.Year)

	if err != nil {
		replyError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	replyJSON(w, r, http.StatusCreated, movie)
}
