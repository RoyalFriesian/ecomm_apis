package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ecomm "github.com/ecomm_apis/ecomm/products"
	"github.com/gorilla/mux"
)

const (
	dbUser = ""
	dbPass = ""
	dbName = "storeDB"
)

func main() {
	//updated 2
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(dbName)

	// create a repository
	repository := ecomm.NewProduct(ctx, collection)

	fmt.Println("Main - Entry")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1.0/product/{id}", repository.GetProduct).Methods("GET")
	r.HandleFunc("/api/v1.0/product", ecomm.AddProduct).Methods("POST")
	r.HandleFunc("/api/v1.0/product", ecomm.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/v1.0/product", ecomm.PutProduct).Methods("PUT")
	r.HandleFunc("/api/v1.0/product", ecomm.PatchProduct).Methods("PATCH")
	http.ListenAndServe(":80", r)
}
