package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TimmyBen/cyoa"
)

func main() {
	//PARSING THE JSON
	// Fetch flag. parse it and print it and then test (filename)

	port := flag.Int("port", 3000, "Port to start CYOA web on")
	filename := flag.String("file", "gopher.json", "The JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open file and handle error
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

// Add a Port flag in main.go with default 3000 and port string "port to start CYOA web on"

// pass in the story to cyoa.NewHandler and get a handler back
// Print out "Starting the server at: port"
// Use http.ListenAndServe but wrap in log.Fatal
