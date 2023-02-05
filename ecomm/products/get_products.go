package ecomm

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "<h1>You've requested the book: %s on page %s\n</h1>", title, page)
}
