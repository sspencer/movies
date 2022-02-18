package services

import (
	"fmt"
	"movies/web"
	"time"
)

type MovieService struct {
	movies []Movie
}

type Movie struct {
	ID       int                `json:"id"`
	Title    string             `json:"title"`
	Year     int                `json:"year"`
	Search   []web.SearchResult `json:"search,omitempty"`
	Duration string             `json:"duration,omitempty"`
}

func NewMovieService() *MovieService {
	movies := []Movie{
		{ID: 1, Title: "The Shawshank Redemption", Year: 1994},
		{ID: 2, Title: "The Godfather", Year: 1972},
		{ID: 3, Title: "The Godfather: Part II", Year: 1974},
		{ID: 4, Title: "The Dark Knight", Year: 2008},
		{ID: 5, Title: "12 Angry Men", Year: 1957},
	}

	return &MovieService{movies}
}

func (s *MovieService) GetMovies() *[]Movie {
	return &s.movies
}

func (s *MovieService) AddMovie(title string, year int) (*Movie, error) {
	movie := Movie{
		ID:    len(s.movies) + 1,
		Title: title,
		Year:  year,
	}

	movie.ID = len(s.movies) + 1
	s.movies = append(s.movies, movie)

	return &movie, nil
}

func (s *MovieService) GetMovie(id int) (*Movie, error) {
	if id > 0 && id <= len(s.movies) {
		movie := s.movies[id-1]
		searchResults, duration := runSearch(web.SerialSearch, movie.Title)
		movie.Search = searchResults
		movie.Duration = duration
		return &movie, nil
	}

	return nil, fmt.Errorf("movie %d not found", id)
}

func runSearch(search web.SearchAlgorithm, title string) ([]web.SearchResult, string) {
	start := time.Now()
	results := search(title)
	elapsed := time.Since(start)
	return results, elapsed.String()
}
