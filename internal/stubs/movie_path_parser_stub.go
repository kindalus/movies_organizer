package stubs

import "kindalus/movies_organizer/internal/organizer"

type moviePathParserStub struct{}

func NewMoviePathParser() organizer.MoviePathParser {
	return new(moviePathParserStub)
}

func (p *moviePathParserStub) Parse(path string) (string, uint16) {
	return "Rambo", 1987
}
