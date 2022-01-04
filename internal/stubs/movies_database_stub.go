package stubs

import (
	"kindalus/movies_organizer/internal/organizer"
)

type moviesDatabaseStub struct {
}

func NewMoviesDatabase() organizer.MoviesDatabase {
	return new(moviesDatabaseStub)
}

func (m *moviesDatabaseStub) Find(name string, year uint16) (*organizer.MovieSpec, error) {

	return &organizer.MovieSpec{
		Title: "Save Private Ryan",
		Year:  "1998",
		Genre: "Action",
	}, nil
}
