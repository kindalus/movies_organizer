package strategies

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kindalus/movies_organizer/internal/organizer"
	"kindalus/movies_organizer/pkg/coalesce"
	"net/http"
	"strings"

	_ "embed"
)

const omdbUrl = "https://www.omdbapi.com/?apikey="

type omdb struct {
	baseUrl string
}

func NewOmdb(key string) (organizer.MoviesDatabase, error) {

	return &omdb{omdbUrl + key}, nil
}

func (db *omdb) Find(title string, year uint16) (*organizer.MovieSpec, error) {
	imdbId, err := db.findMovieImdbId(title, year)
	if err != nil {
		return nil, err
	}

	return db.getMovieSpec(imdbId)

}

func (db *omdb) getMovieSpec(imdbId string) (*organizer.MovieSpec, error) {
	// https://www.omdbapi.com/?apikey=4be87309i=tt0133093
	url := fmt.Sprintf("%s&i=%s", db.baseUrl, imdbId)
	movieResponse := new(omdbMovieResponse)

	if err := getOmdbUrl(url, movieResponse); err != nil {
		return nil, err
	}

	movieSpec := new(organizer.MovieSpec)

	movieSpec.Genre = strings.Split(movieResponse.Genre, ",")[0]
	movieSpec.Title = movieResponse.Title
	movieSpec.Year = movieResponse.Year

	return movieSpec, nil
}

func (db *omdb) findMovieImdbId(title string, year uint16) (string, error) {
	const nilString = ""

	scapedTitle := strings.ReplaceAll(title, " ", "+")

	// https://www.omdbapi.com/?apikey=4be87309&s=Matrix&y=1999
	url := fmt.Sprintf("%s&s=%s&y=%d", db.baseUrl, scapedTitle, year)

	result := new(omdbSearchResponse)
	if err := getOmdbUrl(url, result); err != nil {
		return nilString, err
	}

	return result.Search[0].ImdbID, nil
}

func getOmdbUrl(url string, value interface{}) error {
	resp, errGet := http.Get(url)
	if errGet != nil {
		return errGet
	}

	data, errData := io.ReadAll(resp.Body)
	if errData != nil {
		return errData
	}

	response := new(omdbResponse)
	if err := json.Unmarshal(data, response); err != nil || response.Response != "True" {
		return coalesce.Coalesce(err, errors.New(response.Error))
	}

	if err := json.Unmarshal(data, value); err != nil {
		return err
	}

	return nil
}

type omdbSearchResponse struct {
	Search []omdbSearchResponseMovie
}

type omdbSearchResponseMovie struct {
	Title  string
	Year   string
	ImdbID string
}

type omdbMovieResponse struct {
	Title string
	Year  string
	Genre string
}

type omdbResponse struct {
	Response string
	Error    string
}
