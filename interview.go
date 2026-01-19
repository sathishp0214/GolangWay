package main

// Implement the handler functions GetKeys, GetKey, CreateKey, and DeleteKey. Use
// in-memory storage; the keys do not need to be persisted across restart. The key
// structure is given. How you manage the storage is up to you.
//
// The get methods should only return keys that are unexpired. If the current
// time is after the key's expiration, it should not be returned in the
// response.
//
// All handler methods should return json data.
//
// Example requests and responses:
//
// # Get all unexpired keys.
// curl --verbose --request GET \
//		--url 'http://localhost:8080/keys'
// -> 200
//		[{"id":"<uuid>","expires":"<iso date>"},{"id":"<uuid>","expires":"<iso date>"}]
//
// # Get a specific key. If the key is expired, the response should be Not Found/404.
// curl --verbose --request GET \
//		--url 'http://localhost:8080/keys/<uuid>'
// -> 200
//		{"id":"<uuid>","expires":"<iso date>"}
// -> 404
//
// # Create a key with no expiration. The expiration should be one hour from
// # when the key is created.
// curl --verbose --request POST \
//		--url 'http://localhost:8080/keys' \
//		--header 'Content-type: application/json' \
//		--data '{}'
// -> 204
//		{"id":"<uuid>","expires":"<iso date now + 1 hour>"}
//
// # Create a key with explicit expiration. An expires date in the past is accepted.
// curl --verbose --request POST \
//		--url 'http://localhost:8080/keys' \
//		--header 'Content-type: application/json' \
//		--data '{"expires":"<iso date>"}'
// -> 204
//		{"id":"<uuid>","expires":"<iso date>"}
//
// # Delete a specific key.
// curl --verbose --request DELETE \
//		--url 'http://localhost:8080/keys/<uuid>'
// -> 204
// -> 404
//

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Key struct {
	Id      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

var keysMap = map[string]string{}

func main() {
	router := mux.NewRouter()
	sub := router.PathPrefix("/keys").Subrouter()
	sub.HandleFunc("", GetKeys).
		Methods("GET")
	sub.HandleFunc("/{key_id}", GetKey).
		Methods("GET")
	sub.HandleFunc("", CreateKey).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/json")
	sub.HandleFunc("/{key_id}", DeleteKey).
		Methods("DELETE")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

// Get all keys that are not expired. This should return a JSON array of keys.
func GetKeys(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" || r.Method != "GET" {
		_, err := w.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02 15:04:05"

	filterMap := map[string]string{}

	now := time.Now()
	for i, j := range keysMap {
		date, _ := time.Parse(layout, j)
		if date.Sub(now).Hours() < 1 {
			filterMap[i] = j
		}
	}
	responsedata, err := json.Marshal(filterMap)

	if err != nil {
		panic(err)
	}

	_, err = w.Write([]byte(responsedata))
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		panic(err)
	}
	log.Println("GetKeys")
}

// Get a key by id. If the key is expired, return 404. This should return a JSON object.
func GetKey(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" || r.Method != "GET" {
		_, err := w.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := Key{}
	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		panic(err)
	}

	if keysMap[key.Id] == "" {
		w.WriteHeader(http.StatusNoContent)

		if err != nil {
			panic(err)
		}
		_, err = w.Write([]byte("{}"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseBody := fmt.Sprintf(`{"id": "%v", "expires": "%v"}`, key.Id, keysMap[key.Id])

	_, err = w.Write([]byte(responseBody))
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("GetKey")
}

// Create a new key with the specified expiration. If no expiration is provided,
// use the default of 1 hour. Return the new key as the response.
// Example: {"expires": "2019-01-01T12:00:00Z"}
func CreateKey(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" || r.Method != "POST" {
		_, err := w.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := Key{}
	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		panic(err)
	}

	if key.Expires.GoString() != "" {
		uuid := uuid.New().String()
		keysMap[uuid] = key.Expires.GoString()
		responseBody := fmt.Sprintf(`{"id": "%v", "expires": "%v"}`, uuid, keysMap)
		_, err := w.Write([]byte(responseBody))
		if err != nil {
			panic(err)
		}

		return
	}

	uuid := uuid.New().String()

	expirationTime := time.Now().Add(-1 * time.Hour)
	responseBody := fmt.Sprintf(`{"id": "%v", "expires": "%v"}`, uuid, expirationTime)
	keysMap[uuid] = expirationTime.GoString()
	_, err = w.Write([]byte(responseBody))
	if err != nil {
		panic(err)
	}

	return

}

// Delete the specified key. If the key is expired, return 404.
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" || r.Method != "DELETE" {
		_, err := w.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := Key{}
	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		panic(err)
	}

	if keysMap[key.Id] == "" {
		w.WriteHeader(http.StatusNoContent)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
		return
	}

	delete(keysMap, key.Id)
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		panic(err)
	}
	log.Println("DeleteKey")
}
