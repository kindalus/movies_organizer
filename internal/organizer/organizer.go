package organizer

import "path"

type OrganizerContext struct {
	StorageProvider StorageProvider
	MoviesDatabase  MoviesDatabase
	MoviePathParser MoviePathParser
}

type Organizer struct {
	OrganizerContext
	destinationPath string
}

func New(ctx OrganizerContext, destinationPath string) (*Organizer, error) {

	o := &Organizer{ctx, destinationPath}

	if err := o.verifyDir(destinationPath, ErrDestinationPathNotFound); err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Organizer) Organize(moviePath string) (string, error) {

	if moviePath == "" {
		return "", ErrNoMovieGiven
	}

	if err := o.verifyDir(moviePath, ErrMoviePathNotFound); err != nil {
		return "", err
	}

	name, year := o.MoviePathParser.Parse(moviePath)

	spec, err := o.MoviesDatabase.Find(name, year)
	if err != nil {
		return "", err
	}

	destinationPath := path.Join(o.destinationPath, spec.Year, spec.Genre)
	if err := o.createFolderIfDoesntExist(destinationPath); err != nil {
		return "", err
	}

	o.StorageProvider.Move(moviePath, destinationPath)

	outputPath := path.Join(destinationPath, path.Base(moviePath))

	return outputPath, nil
}

func (o *Organizer) createFolderIfDoesntExist(destinationPath string) error {
	destinationExists, err := o.StorageProvider.DirExists(destinationPath)
	if err != nil {
		return err
	}

	if destinationExists {
		return nil
	}

	return o.StorageProvider.MkDir(destinationPath)
}

func (o *Organizer) verifyDir(path string, onError error) error {
	exists, err := o.StorageProvider.DirExists(path)

	if err != nil {
		return err
	}

	if !exists {
		return onError
	}

	return nil
}
