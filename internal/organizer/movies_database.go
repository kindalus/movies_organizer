package organizer

type MovieSpec struct {
	Title string
	Year  string
	Genre string
}

type MoviesDatabase interface {
	Find(name string, year uint16) (*MovieSpec, error)
}
