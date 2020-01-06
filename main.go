package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var client *redis.Client

//Name is something
type Name struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
}

// var names []Name

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	r := mux.NewRouter()

	// names = append(names, Name{ID: "0", Firstname: "John"})
	// names = append(names, Name{ID: "1", Firstname: "Kavish"})

	r.HandleFunc("/names", createName).Methods("POST")
	r.HandleFunc("/names", getNames).Methods("GET")
	r.HandleFunc("/names/{id}", getName).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func getNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	val, err := client.LRange("names", 0, -1).Result()
	log.Println(val, len(val))
	check(err)
	json.NewEncoder(w).Encode(client.LRange("names", 0, -1))
}

func getName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	value := client.LIndex("names", id)
	log.Println(value.Result())
	check(value.Err())
	// for _, item := range names {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Name{})
}

func createName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name Name
	_ = json.NewDecoder(r.Body).Decode(&name)
	err := client.RPush("names", name.Firstname).Err()
	if err != nil {
		log.Println("error inserting redis", err)
	}
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(names)
}

// func indexSetHandler(w http.ResponseWriter, r *http.Request) {
// 	err := client.Set("key", "value", 0).Err()
// 	check(err)
// }

// func indexGetHandler(w http.ResponseWriter, r *http.Request) {
// 	val, err := client.Get("key").Result()
// 	check(err)
// 	fmt.Println("key", val)
// }

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
