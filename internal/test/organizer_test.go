package organizer

import (
	"kindalus/movies_organizer/internal/mocks"
	"kindalus/movies_organizer/internal/organizer"
	"kindalus/movies_organizer/internal/stubs"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Recebe a directoria do filme e o destino
// Devolve a directoria onde foi arrumada ou error

func TestOrganizer_Organize(t *testing.T) {
	t.Run("Devolve erro @ErrDestinationPathNotFound se a directoria de destino não existir", func(t *testing.T) {
		destPath := "/tmp/movies"
		storageProvider := mocks.NewStorageProvider()
		storageProvider.On("DirExists", destPath).Return(false, nil)

		_, err := newOrganizer(organizer.OrganizerContext{StorageProvider: storageProvider}, destPath)

		storageProvider.AssertExpectations(t)
		assert.Equal(t, organizer.ErrDestinationPathNotFound, err)

	})

	t.Run("Devolve erro @ErrMoviePathNotFound se a directoria do filme não existir", func(t *testing.T) {
		destPath := "/tmp/movies"
		moviePath := "/movies/Santana.2021.1080p"
		storageProvider := mocks.NewStorageProvider()
		storageProvider.On("DirExists", destPath).Return(true, nil)
		storageProvider.On("DirExists", moviePath).Return(false, nil)

		moviesOrganizer, _ := organizer.New(organizer.OrganizerContext{StorageProvider: storageProvider}, destPath)

		_, err := moviesOrganizer.Organize(moviePath)

		storageProvider.AssertExpectations(t)
		assert.Equal(t, organizer.ErrMoviePathNotFound, err)
	})

	t.Run("Devolve erro @ErrNoMovieGiven se o filme não for especificado", func(t *testing.T) {
		moviesOrganizer, _ := newOrganizer(organizer.OrganizerContext{}, "")

		_, err := moviesOrganizer.Organize("")

		assert.Equal(t, organizer.ErrNoMovieGiven, err)
	})

	t.Run("Constroi a directoria de destino com o formato [Base Dir]_[Ano]_[Genero]_[Movie Path]", func(t *testing.T) {
		destPath := "/downloads/movies"
		moviePath := "/movies/The.Good.Movie.2021.1080p"
		expectedOutput := path.Join(destPath, "2020", "Animation", "The.Good.Movie.2021.1080p")

		mdb := mocks.NewMoviesDatabase()
		mdb.On("Find", mock.AnythingOfType("string")).
			Return(&organizer.MovieSpec{Title: "The Good Movie", Genre: "Animation", Year: "2020"}, nil)

		moviesOrganizer, _ := newOrganizer(organizer.OrganizerContext{MoviesDatabase: mdb}, destPath)

		output, err := moviesOrganizer.Organize(moviePath)

		mdb.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedOutput, output)
	})

}

func newOrganizer(ctx organizer.OrganizerContext, basePath string) (*organizer.Organizer, error) {
	return organizer.New(
		organizer.OrganizerContext{
			StorageProvider: coalesce(ctx.StorageProvider, stubs.NewStorageProvider()),
			MoviesDatabase:  coalesce(ctx.MoviesDatabase, stubs.NewMoviesDatabase()),
			MoviePathParser: coalesce(ctx.MoviePathParser, stubs.NewMoviePathParser()),
		},
		basePath)
}

func coalesce[T any](a T, b T) T {
	if interface{}(a) == nil {
		return b
	}

	return a
}
