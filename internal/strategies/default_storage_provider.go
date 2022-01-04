package strategies

import (
	"kindalus/movies_organizer/internal/organizer"
	"os"
	"path"
)

type defaultStorageProvider struct{}

func NewDefaultStorageProvider() organizer.StorageProvider {
	return new(defaultStorageProvider)
}

func (p *defaultStorageProvider) DirExists(path string) (bool, error) {
	_, err := os.Open(path)

	if err != nil && os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *defaultStorageProvider) Mkdir(path string) error {
	return os.MkdirAll(path, os.ModeDir|os.ModePerm)
}

func (p *defaultStorageProvider) Move(src, dest string) error {
	return os.Rename(src, path.Join(dest, path.Base(src)))
}
