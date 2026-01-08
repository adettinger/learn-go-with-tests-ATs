package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.URL.Query().Get("name"))
}
