package main

import (
	_ "embed"
	"flag"
	"fmt"
	"kindalus/movies_organizer/internal/organizer"
	"kindalus/movies_organizer/internal/strategies"
	"os"
)

//go:embed omdb.key
var omdbKey string

func main() {

	dry := flag.Bool("dry", false, "Dry Run")

	flag.Parse()
	args := clearArgs()

	if len(args) < 2 {
		fmt.Println("expected [-dry] Dest ...[Movie Path]")
		os.Exit(1)
	}

	context := organizer.OrganizerContext{
		StorageProvider: strategies.NewDefaultStorageProvider(),
		MoviesDatabase:  strategies.NewOmdb(omdbKey),
		MoviePathParser: strategies.NewDefaultMoviePathParser(),
	}

	if *dry {
		context.StorageProvider = strategies.NewDryStorageProvider()
	}

	organizer, err := organizer.New(context, args[0])

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for _, movie := range args[1:] {
		fmt.Println("\nMovie:", movie)

		result, err := organizer.Organize(movie)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result:", result)
	}
}

func clearArgs() []string {
	result := make([]string, 0)

	for _, arg := range os.Args[1:] {
		if arg != "-dry" {
			result = append(result, arg)
		}

	}

	return result
}
