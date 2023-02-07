package main

import (
	"fmt"
	"net/http"

	ecomm "github.com/ecomm_apis/ecomm/products"
	"github.com/gorilla/mux"
)

func main() {
	//updated 2

	fmt.Println("Main - Entry")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1.0/product/{name}", ecomm.GetProduct).Methods("GET")
	r.HandleFunc("/api/v1.0/product", ecomm.AddProduct).Methods("POST")
	r.HandleFunc("/api/v1.0/product", ecomm.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/v1.0/product", ecomm.PutProduct).Methods("PUT")
	r.HandleFunc("/api/v1.0/product", ecomm.PatchProduct).Methods("PATCH")
	http.ListenAndServe(":80", r)
}
