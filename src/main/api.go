package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway" // <- gateway to AWS lambda functions
	"github.com/julienschmidt/httprouter"
	"io.github.csarnataro/swapi-go/src/film"
	"io.github.csarnataro/swapi-go/src/films"
)

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
	}

	router := httprouter.New()

	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
	})

	router.GET("/api/films", films.Handler)
	router.GET("/api/films/:id", film.Handler)

	fmt.Printf("Server listening on port %d...\n", *port)
	log.Fatal(listener(portStr, router))
}
