package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	if request.Data == nil {
		http.Error(rw, "missing field 'data' from JSON object", http.StatusBadRequest)
		return
	}

	if request.DataPoints == nil {
		http.Error(rw, "missing field 'points' from JSON object", http.StatusBadRequest)
		return
	}

	if decoder.More() {
		http.Error(rw, "extraneous data after JSON object", http.StatusBadRequest)
		return
	}

	log.Println(*request.Data)
	log.Println(*request.DataPoints)
}

func main() {
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("Unable to start web server", err)
	}
}
