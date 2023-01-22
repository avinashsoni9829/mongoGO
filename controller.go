package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserId string `json:"_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

var collection = dbConnection().Database("mongoGo").Collection("users")

// create Users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	response, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserting a new document")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "new user added %s\n", response)

}

// findUserDetailsById

func GetUserDetailsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userId"]
	fmt.Println("searching for user with id = " + id)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	var result primitive.M
	resp := collection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&result)

	if resp != nil {
		log.Fatal(resp)

	}

	json.NewEncoder(w).Encode(result)
	fmt.Println("inserting a new document")
	w.WriteHeader(http.StatusFound)

	fmt.Fprintf(w, "getting User Details %s\n", result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	e := json.NewDecoder(r.Body).Decode(&user)

	if e != nil {
		log.Fatal(e)
	}

	filter := bson.D{{"userid", user.UserId}}
	after := options.After

	returnVal := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{{"$set", bson.D{{
		"name", user.Name}, {"age", user.Age}}}}

	updateRes := collection.FindOneAndUpdate(context.TODO(), filter, update, &returnVal)

	var result primitive.M

	_ = updateRes.Decode(&result)

	json.NewEncoder(w).Encode(result)
	fmt.Printf("document Updated Successfully")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userId"]
	fmt.Println("deleting user with id = " + id)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	opts := options.Delete().SetCollation(&options.Collation{})

	response, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleting document ")

	json.NewEncoder(w).Encode(response.DeletedCount)

}
