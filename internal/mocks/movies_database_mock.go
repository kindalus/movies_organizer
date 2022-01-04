package mocks

import (
	"kindalus/movies_organizer/internal/organizer"

	"github.com/stretchr/testify/mock"
)

type moviesDatabaseMock struct {
	mock.Mock
}

func NewMoviesDatabase() *moviesDatabaseMock {
	return new(moviesDatabaseMock)
}

func (m *moviesDatabaseMock) Find(name string, year uint16) (*organizer.MovieSpec, error) {
	args := m.Called(name)

	return args.Get(0).(*organizer.MovieSpec), args.Error(1)
}
