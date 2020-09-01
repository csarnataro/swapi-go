package films

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/csarnataro/swapi-go/src/constants"
)

//go:generate go run gen.go

func sendNotFoundError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, constants.NotFoundJSON)
}

// Handler returns the full list of films
func Handler(w http.ResponseWriter, r *http.Request) { // , params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	// fmt.Fprintf(w, `{"result": "ok"}`)
	var pageNumber uint64 = 1
	var conversionError error = nil
	page := r.URL.Query().Get("page") // .params.ByName("page")
	if page != "" {
		pageNumber, conversionError = strconv.ParseUint(page, 10, 0)
		if conversionError != nil {
			sendNotFoundError(w)
			return
		}
	} else {
		pageNumber = 1
	}

	fmt.Println("Requested page number:", pageNumber)

	content := Films

	var entries []FilmEntry
	// parsing JSON file
	err := json.Unmarshal([]byte(content), &entries)
	if err != nil {
		log.Fatal(fmt.Println("error:", err))
	}
	result, err := buildResult(entries, pageNumber)

	if err != nil {
		sendNotFoundError(w)
		return
	}
	// firstFilm := originalJSON[0]
	destJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "%s", destJSON)
}
