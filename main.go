package main

import (
	"fmt"
	"net/http"

	ecomm "github.com/ecomm_apis/ecomm/products"
	"github.com/gorilla/mux"
)

func main() {
	//updated
	fmt.Println("Main - Entry")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1.0/product/{name}", ecomm.GetProduct).Methods("GET")
	http.ListenAndServe(":80", r)
}
