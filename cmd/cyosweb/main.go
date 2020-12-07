package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/impzero/cyos"
)

func main() {
	port := flag.Int("port", 8080, "The port to run the web server on")
	filename := flag.String("file", "story.json", "The file that contains the story structure in JSON")
	flag.Parse()
	fmt.Printf("Using the story in %s!\n", *filename)

	// Open the file
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyos.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyos.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	fmt.Printf("%+v\n", story)
}
