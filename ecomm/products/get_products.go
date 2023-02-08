package ecomm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type productCtx struct {
	db   *mongo.Database
	pCtx context.Context
}

func NewProduct(ctx context.Context, db *mongo.Database) *productCtx {
	return &productCtx{db: db, pCtx: ctx}
}

type Product struct {
	Product_ID   int     `json:"product_id" bson:"product_id"`
	Product_Name string  `json:"product_name" bson:"product_name"`
	Retail_Price float64 `json:"retail_price" bson:"retail_price"`
}

type Status struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (pc *productCtx) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		status := Status{ID: 0, Status: "Failure:" + err.Error()}
		res, err := json.Marshal(status)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	var product Product
	err = pc.db.
		Collection("products").
		FindOne(pc.pCtx, bson.M{"product_id": product_id}).
		Decode(&product)
	if err != nil {
		status := Status{ID: 0, Status: "Failure:" + err.Error()}
		res, err := json.Marshal(status)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}

	res, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (pc *productCtx) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := pc.db.Collection("products").InsertOne(pc.pCtx, product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Inserted ID : ", result.InsertedID)
	res, err := json.Marshal(product)

	if err != nil {
		fmt.Println("Error : ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
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
	var deleteStatus Status
	deleteStatus = Status{ID: product.Product_ID, Status: "Success : Product Deleted"}
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
	var putStatus Status
	putStatus = Status{ID: product.Product_ID, Status: "Success : Product Updated"}
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
	var patchStatus Status
	patchStatus = Status{ID: product.Product_ID, Status: "Success : Product Patch Updated"}
	res, err := json.Marshal(patchStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
