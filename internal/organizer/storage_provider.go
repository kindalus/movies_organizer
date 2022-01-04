package organizer

type StorageProvider interface {
	DirExists(path string) (bool, error)
	Move(source string, destination string) error
	Mkdir(path string) error
}
