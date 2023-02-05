package main

import (
	"fmt"
	"net/http"

	ecomm "github.com/ecomm_apis/ecomm/products"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Main - Entry")
	r := mux.NewRouter()
	r.HandleFunc("api/v1.0/product/{name}", ecomm.CreateBook).Methods("GET")
	r.HandleFunc("api/v1.0/products/", ecomm.CreateBook).Methods("POST")
	r.HandleFunc("api/v1.0/products/{name}", ecomm.CreateBook).Methods("PUT")
	r.HandleFunc("api/v1.0/products/{title}", ecomm.CreateBook).Methods("DELETE")
	http.ListenAndServe(":80", r)
}
