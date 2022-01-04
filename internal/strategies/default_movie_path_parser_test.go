package strategies

import (
	"fmt"
	"testing"
)

func TestDefaultMoviePath_Parse(t *testing.T) {
	parser := NewDefaultMoviePathParser()

	for _, testString := range getMovies() {
		title, year := parser.Parse(testString)

		fmt.Printf("File: %s\nTitle: %s\nYear: %d\n\n", testString, title, year)
	}
}

func getMovies() []string {

	return []string{
		"		Mundo.Dos.Sabios.2021.720p",
		"		/tmp/Mundo.Dos.Sabios.2021.720p",
		"/tmp/   Mundo.Dos.Sabios.2021.720p",
		"/tmp/movies/Last.Looks.2021.1080p.WEBRip.x264-RARBG",
		"Queenpins.2021.1080p.AMZN.WEBRip.DDP5.1.x264-NOGRP",
		"Between Two Ferns The Movie (2019) [WEBRip] [1080p] [YTS.LT]",
		"Otherhood (2019) [WEBRip] [1080p] [YTS.LT]",
		"Jexi (2019) [WEBRip] [1080p] [YTS.LT]",
		"/movies/Parasite (2019) [BluRay] [1080p] [YTS.LT]",
		"Jojo.Rabbit.2019.1080p.BluRay.x264-YOL0W[rarbg]",
		"Playing.with.Fire.2019.1080p.BluRay.x264-AAA[rarbg]",
		"Knives.Out.2019.1080p.WEBRip.DD5.1.x264-CM",
		"/2012/Plus One (2019) [WEBRip] [1080p] [YTS.LT]",
		"Late Night (2019) [WEBRip] [1080p] [YTS.LT]",
		"Tall Girl (2019) [WEBRip] [1080p] [YTS.LT]",
		"Little (2019) [WEBRip] [1080p] [YTS.LT]",
		"The Hustle (2019) [WEBRip] [1080p] [YTS.LT]",
		"Long Shot (2019) [WEBRip] [720p] [YTS.LT]",
		"The.Beach.Bum.2019.1080p.BluRay.x264-DRONES[rarbg]",
		"Long.Shot.2019.1080p.WEB-DL.DD5.1.H264-FGT",
		"The.Gentlemen.2019.1080p.WEBRip.x265-RARBG",
		"Marriage.Story.2019.1080p.WEBRip.x264-RARBG     Vivarium.2019.1080p.WEBRip.x265-RARBG",
		"Once Upon A Time ... In Hollywood (2019) [WEBRip] [1080p] [YTS.LT]",
		"Yesterday (2019) [BluRay] [720p] [YTS.LT]",
		"2032 (2019)",
		"Seal.Team.2021.1080p.WEBRip.x265-RARBG",
		"S.E.A.L.Team.2021.1080p.WEBRip.x265-RARBG",
	}
}
