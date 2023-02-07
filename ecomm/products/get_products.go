package ecomm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Details  string   `json:"details"`
	Variants []string `json:"variants"`
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	p := Product{ID: 10, Name: name, Details: "Some details"}
	res, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	product_name := r.FormValue("product_name")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "<h1>Requesting to add the product: %s \n</h1>", product_name)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	product_id := r.FormValue("product_id")
	fmt.Fprintf(w, "<h1>Deleting product : %s </h1>\n", product_id)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	product_id := r.FormValue("product_id")
	fmt.Fprintf(w, "<h1> Updating Product : %s  Done!</h1>\n", product_id)
}

func PatchProduct(w http.ResponseWriter, r *http.Request) {
	product_id := r.FormValue("product_id")
	fmt.Fprintf(w, "<h1>Product Updated : %s Done!</h1>\n", product_id)
}
