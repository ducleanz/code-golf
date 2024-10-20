package main

import (
	"code_golf/pkg/data_path_parser"
	"flag"
	"fmt"
	"log"
)

func main() {
	feature := flag.String("feature", "", "The feature to run (feature1, feature2)")
	path := flag.String("path", "", "The path to parse")
	flag.Parse()

	switch *feature {
	case "data_path_parser":
		if *path == "" {
			fmt.Println("Please provide a path to parse using the -path flag.")
			return
		}
		segments, err := data_path_parser.ParsePath(*path)
		if err != nil {
			log.Printf("Error parsing path %q: %v\n", *path, err)
		} else {
			fmt.Printf("Path: %q -> Segments: %v\n", *path, segments)
		}
	default:
		fmt.Println("Please specify a valid feature to run using the -feature flag.")
	}
}
