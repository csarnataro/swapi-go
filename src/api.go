package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway" // <- gateway to AWS lambda functions
	people "github.com/csarnataro/swapi-go/src/people/utils"
	films "github.com/csarnataro/swapi-go/src/films/utils"
)

//go:generate go run ./generator/generate-models.go films
//go:generate go run ./generator/generate-models.go people
//go:generate go run ./generator/generate-models.go planets
//go:generate go run ./generator/generate-models.go species
//go:generate go run ./generator/generate-models.go starships
//go:generate go run ./generator/generate-models.go transport
//go:generate go run ./generator/generate-models.go vehicles

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
	}

	http.HandleFunc("/api/people", people.Handler)
	http.HandleFunc("/api/films", films.Handler)

	fmt.Printf("Server listening on port %d...\n", *port)
	log.Fatal(listener(portStr, nil))
}
