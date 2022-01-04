package organizer

type MoviePathParser interface {
	Parse(path string) (string, uint16)
}
