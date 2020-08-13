package film

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler returns a single film
func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "A single film here!\n")
}
