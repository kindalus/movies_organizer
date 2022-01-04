package stubs

import "kindalus/movies_organizer/internal/organizer"

type storageProviderStub struct {
}

func NewStorageProvider() organizer.StorageProvider {
	return new(storageProviderStub)
}

func (m *storageProviderStub) DirExists(path string) (bool, error) {
	return true, nil
}

func (m *storageProviderStub) Move(source string, destination string) error {
	return nil
}

func (m *storageProviderStub) Mkdir(path string) error {
	return nil
}
