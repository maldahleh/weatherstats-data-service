package main

import (
	"encoding/json"
	"log"
	"net/http"

	"weatherstatsData/handlers"
	data "weatherstatsData/request"
)

func handleRequest(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	var request data.DataRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if !request.Validate() {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	if decoder.More() {
		http.Error(rw, "extraneous data after JSON object", http.StatusBadRequest)
		return
	}

	for key, value := range *request.Data {
		handlers.RetrieveData(key, value, request.DataPoints)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("Unable to start web server", err)
	}
}
