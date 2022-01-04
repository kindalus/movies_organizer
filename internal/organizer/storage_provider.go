package organizer

type StorageProvider interface {
	DirExists(path string) (bool, error)
	Move(source string, destination string) error
	MkDir(path string) error
}
