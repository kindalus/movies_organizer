package strategies

import (
	"fmt"
	"kindalus/movies_organizer/internal/organizer"
	"os"
)

type dryStorageProvider struct{}

func NewDryStorageProvider() organizer.StorageProvider {
	return new(dryStorageProvider)
}

func (p *dryStorageProvider) DirExists(path string) (bool, error) {
	_, err := os.Open(path)

	if err != nil && os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *dryStorageProvider) Mkdir(path string) error {
	fmt.Println("mkdir -p ", path)
	return nil
}

func (p *dryStorageProvider) Move(src, dest string) error {
	fmt.Println("mv ", src, dest)
	return nil
}
