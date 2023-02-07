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

type DeleteStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type PutStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type PatchStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
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
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var deleteStatus DeleteStatus
	deleteStatus = DeleteStatus{ID: product.ID, Status: "Success : Product Deleted"}
	res, err := json.Marshal(deleteStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var putStatus PutStatus
	putStatus = PutStatus{ID: product.ID, Status: "Success : Product Updated"}
	res, err := json.Marshal(putStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func PatchProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var patchStatus PatchStatus
	patchStatus = PatchStatus{ID: product.ID, Status: "Success : Product Patch Updated"}
	res, err := json.Marshal(patchStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
