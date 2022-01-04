package mocks

import "github.com/stretchr/testify/mock"

type storageProviderMock struct {
	mock.Mock
}

func NewStorageProvider() *storageProviderMock {
	return new(storageProviderMock)
}

func (m *storageProviderMock) DirExists(path string) (bool, error) {
	args := m.Called(path)

	return args.Bool(0), args.Error(1)
}

func (m *storageProviderMock) Move(source string, destination string) error {
	args := m.Called(source, destination)

	return args.Error(0)
}

func (m *storageProviderMock) MkDir(path string) error {
	args := m.Called(path)

	return args.Error(0)
}
