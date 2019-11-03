package main

import (
	"encoding/json"
	"log"
	"net/http"

	"weatherstatsData/handlers"
	data "weatherstatsData/request"

	logrus "github.com/sirupsen/logrus"
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
		http.Error(rw, "invalid request", http.StatusBadRequest)
		return
	}

	if decoder.More() {
		http.Error(rw, "extraneous data after JSON object", http.StatusBadRequest)
		return
	}

	stationData := make(map[string]handlers.YearlyData)
	for key, value := range *request.Data {
		stationData[key] = handlers.RetrieveData(key, value, request.DataPoints)
	}

	resp, err := json.Marshal(stationData)
	if err != nil {
		logrus.Error("json failed", err)
		http.Error(rw, "request failed, try again later", http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(resp)
	if err != nil {
		logrus.Error("HTTP Write Failure", err)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("Unable to start web server", err)
	}
}
