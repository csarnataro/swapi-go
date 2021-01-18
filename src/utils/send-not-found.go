package utils

import (
	"fmt"
	"net/http"

	"github.com/csarnataro/swapi-go/src/constants"
)

// SendNotFoundError sends a 404 to the specified ResponseWriter
func SendNotFoundError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, constants.NotFoundJSON)
}
