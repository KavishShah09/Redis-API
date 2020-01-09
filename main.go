package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var client *redis.Client

//Object is anything with a key string and a value string
type Object struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	r := mux.NewRouter()

	r.HandleFunc("/data", createValue).Methods("POST")
	r.HandleFunc("/data", getData).Methods("GET")
	r.HandleFunc("/data/{key}", getValue).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Query() == nil {
		var limit int64 = 10
		match := ""
		keys, _, err := client.Scan(0, match, limit).Result()
		if err != nil {
			log.Println("error getting keys", err)
		}
		log.Println(keys)

	} else {
		limitString := r.URL.Query().Get("limit")
		limit, err := strconv.ParseInt(limitString, 10, 64)
		if limit != 0 {
			if err != nil {
				log.Println("error converting string to int64", err)
			}
		}
		match := r.URL.Query().Get("match")
		keys, _, err := client.Scan(0, match, limit).Result()
		if err != nil {
			log.Println("error getting keys", err)
		}
		log.Println(keys)
	}

}

func getValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	val, err := client.Get(params["key"]).Result()
	if err != nil {
		log.Println("error getting value redis", err)
	}
	log.Println(val)
	json.NewEncoder(w).Encode(&val)
}

func createValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var object Object
	_ = json.NewDecoder(r.Body).Decode(&object)
	err := client.Set(object.Key, object.Value, 0).Err()
	if err != nil {
		log.Println("error inserting redis", err)
	}
	w.WriteHeader(http.StatusOK)
}
