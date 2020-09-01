package films

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/csarnataro/swapi-go/src/constants"
)

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

	ex := os.Getenv("LAMBDA_TASK_ROOT")
	if ex == "" {
		ex = "."
	}

	err := filepath.Walk("/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	exPath := path.Join(ex, "functions", "data", "films.json")

	content, err := ioutil.ReadFile(exPath)
	if err != nil {
		fmt.Fprint(w, "Some error occurred: ", err)
	} else {
		var entries []FilmEntry
		// parsing JSON file
		err := json.Unmarshal(content, &entries)
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
}

// func timer(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 		startTime := time.Now()
// 		h.ServeHTTP(w, r)
// 		duration := time.Now().Sub(startTime)

// 	})
// }
