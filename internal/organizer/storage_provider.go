package organizer

type StorageProvider interface {
	DirExists(path string) (bool, error)
}
