package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway" // <- gateway to AWS lambda functions
	film "github.com/csarnataro/swapi-go/src/film/utils"
	// "github.com/julienschmidt/httprouter"
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

	// router := httprouter.New()

	// router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
	// })

	// router.GET("/healthcheck", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	fmt.Fprintf(w, "OK\n")
	// })

	mux := http.NewServeMux()

	mux.HandleFunc("/1", film.Handler)

	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	fmt.Printf("Server listening on port %d...\n", *port)
	log.Fatal(listener(portStr, mux))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/apex/gateway"
// )

// func main() {
// 	http.HandleFunc("/", hello)
// 	log.Fatal(gateway.ListenAndServe(":3000", nil))
// }

func hello(w http.ResponseWriter, r *http.Request) {
	// example retrieving values from the api gateway proxy request context.
	requestContext, ok := gateway.RequestContext(r.Context())
	if !ok || requestContext.Authorizer["sub"] == nil {
		fmt.Fprint(w, "Hello World from Go")
		return
	}

	userID := requestContext.Authorizer["sub"].(string)
	fmt.Fprintf(w, "Hello %s from Go", userID)
}
