package films

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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

	content := Films

	var entries []FilmEntry
	// parsing JSON file
	err := json.Unmarshal([]byte(content), &entries)
	if err != nil {
		log.Fatal(fmt.Println("error:", err))
	}

	// differentiating all films request from single film request

	var allFilms = regexp.MustCompile(`^/(api/)?films\/?$`) // <- /api/ is defined as redirect on netlify
	var singleFilm = regexp.MustCompile(`^/(api/)?films/(\d+)$`)

	path := r.URL.Path
	fmt.Println("Requested URL:", path)

	switch {
	case allFilms.MatchString(path):
		fmt.Println("Requested all films")
		var pageNumber uint64 = 1
		var conversionError error = nil
		page := r.URL.Query().Get("page")
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

		result, err := buildResult(entries, pageNumber)

		if err != nil {
			sendNotFoundError(w)
			return
		}
		destJSON, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "%s", destJSON)
	case singleFilm.MatchString(path):
		ID := strings.TrimPrefix(path, "/films/")
		fmt.Println("Requested single film:", ID)
		for _, film := range entries {
			if strconv.Itoa(film.Pk) == ID {
				result := buildFilm(film)
				destJSON, err := json.Marshal(result)
				if err != nil {
					fmt.Fprintf(w, "Error: %s", err.Error())
					return
				}
				fmt.Fprintf(w, "%s", destJSON)
				return
			}
		}
		sendNotFoundError(w)
	default:
		fmt.Println("Wrong path:", path)
		sendNotFoundError(w)
	}

	// fmt.Fprintf(w, `{"result": "ok"}`)

}
