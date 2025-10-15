package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/TimmyBen/cyoa"
)

func main() {
	//PARSING THE JSON
	// Fetch flag. parse it and print it and then test (filename)
	filename := flag.String("file", "gopher.json", "The JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open file and handle error
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	// create a Json Decoder Reader using io.file reader
	d := json.NewDecoder(f)

	// Decode the story and print it with its fields
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
