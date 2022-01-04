package strategies

import (
	"kindalus/movies_organizer/internal/organizer"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type defaultMoviePathParser struct{}

func NewDefaultMoviePathParser() organizer.MoviePathParser {
	return new(defaultMoviePathParser)
}

func (p *defaultMoviePathParser) Parse(path string) (string, uint16) {
	return toSpec(removePath(cleanup(path)))
}

func cleanup(value string) string {
	return strings.TrimSpace(value)
}

func removePath(value string) string {
	return path.Base(value)
}

func toSpec(value string) (string, uint16) {

	r, _ := regexp.Compile("(.{3,}?)\\W?(\\d{4})\\W?.*")
	groups := r.FindStringSubmatch(value)

	if len(groups) == 0 {
		return "", uint16(0)
	}

	year, _ := strconv.ParseUint(groups[2], 10, 16)
	title := cleanup(removeDots(groups[1]))

	return title, uint16(year)
}

func removeDots(value string) string {
	r, _ := regexp.Compile("(\\w)\\.(\\w{2,})")

	result := r.ReplaceAllString(value, "$1 $2")

	if result == value {
		return result
	}

	return removeDots(result)
}
