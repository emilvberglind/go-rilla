package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message": "Doing Things"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
