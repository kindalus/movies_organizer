package mocks

import (
	"github.com/stretchr/testify/mock"
)

type moviePathParserMock struct {
	mock.Mock
}

func NewMoviePathParser() *moviePathParserMock {
	return new(moviePathParserMock)
}

func (p *moviePathParserMock) Parse(path string) (string, uint16, error) {
	args := p.Called(path)

	return args.String(0), uint16(args.Int(1)), args.Error(2)
}
