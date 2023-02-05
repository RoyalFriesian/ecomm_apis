package ecomm

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "<h1>You've requested the product: %s \n</h1>", name)
}
