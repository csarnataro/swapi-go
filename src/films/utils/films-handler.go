package films

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"regexp"
	"strconv"

	"github.com/csarnataro/swapi-go/src/utils"
	"github.com/csarnataro/swapi-go/generated/films"

)

// Handler returns the full list of films
func Handler(w http.ResponseWriter, r *http.Request) { // , params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	content := generated.Films

	var entries []FilmEntry
	// parsing JSON file
	err := json.Unmarshal([]byte(content), &entries)
	if err != nil {
		log.Fatal(fmt.Println("error:", err))
	}

	// differentiating all films request from single film request

	var allFilms = regexp.MustCompile(`^/api/films\/?$`) // <- /api/ is defined as redirect on netlify
	var singleFilm = regexp.MustCompile(`^/api/films/(\d+)$`)

	requestedPath := r.URL.Path

	switch {
	case allFilms.MatchString(requestedPath):
		fmt.Println("Requested all films")
		var pageNumber uint64 = 1
		var conversionError error = nil
		page := r.URL.Query().Get("page")
		if page != "" {
			pageNumber, conversionError = strconv.ParseUint(page, 10, 0)
			if conversionError != nil {
				utils.SendNotFoundError(w)
				return
			}
		} else {
			pageNumber = 1
		}

		fmt.Println("Requested page number:", pageNumber)

		result, err := buildResult(entries, utils.GetServerName(r), pageNumber)

		if err != nil {
			utils.SendNotFoundError(w)
			return
		}
		destJSON, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "%s", destJSON)
	case singleFilm.MatchString(requestedPath):
		ID := path.Base(requestedPath)
		fmt.Println("Requested single film:", ID)
		for _, film := range entries {
			if strconv.Itoa(film.Pk) == ID {
				result := buildFilm(film, utils.GetServerName(r))
				destJSON, err := json.Marshal(result)
				if err != nil {
					fmt.Fprintf(w, "Error: %s", err.Error())
					return
				}
				fmt.Fprintf(w, "%s", destJSON)
				return
			}
		}
		utils.SendNotFoundError(w)
	default:
		fmt.Println("Wrong path:", requestedPath)
		utils.SendNotFoundError(w)
	}

	// fmt.Fprintf(w, `{"result": "ok"}`)

}
