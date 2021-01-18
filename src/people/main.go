package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway" // <- gateway to AWS lambda functions
	people "github.com/csarnataro/swapi-go/src/people/utils"
)

//go:generate go run ../generator/generate-models.go people

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
	}

	http.HandleFunc("/", people.Handler)

	fmt.Printf("Server listening on port %d...\n", *port)
	log.Fatal(listener(portStr, nil))
}
