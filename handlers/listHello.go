package handlers

import (
	"fmt"
	"net/http"
)

func ListHello(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "ichnaea server is live!")
	}
}