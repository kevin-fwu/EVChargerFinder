package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReqParams struct {
	Latitude   float64
	Longitude  float64
	Distance   float64
	CountLimit int
}

func httpLookup(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var parms ReqParams
	err := decoder.Decode(&parms)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to decode input JSON. Error: %+v", err)
		return
	}

	list := findClosest(&parms)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(list)
}

func listen(address, sslcert, sslkey string) {

	http.HandleFunc("/getnearest", httpLookup)

	if sslkey != "" {
		log.Fatal(http.ListenAndServeTLS(address, sslcert, sslkey, nil))
	} else {
		log.Fatal(http.ListenAndServe(address, nil))
	}
}
