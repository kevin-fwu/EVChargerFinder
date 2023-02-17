package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func httpLookup(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to parse input body. Error: %+v", err)
		return
	}

	arg, ok := r.PostForm["coordinates"]
	if !ok || len(arg) != 1 {
		fmt.Fprintf(w, "Invalid input coordinate count")
		return
	}

	var latitude, longitude float64

	_, err = fmt.Sscanf(arg[0], "%f,%f", &latitude, &longitude)

	if err != nil {
		fmt.Fprintf(w, "Invalid input coordinates")
		return
	}

	arg, ok = r.PostForm["distance"]
	if !ok || len(arg) != 1 {
		fmt.Fprintf(w, "Invalid input distance count")
	}

	dist, err := strconv.ParseFloat(arg[0], 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid input distance")
		return
	}

	arg, ok = r.PostForm["limit"]
	if !ok || len(arg) != 1 {
		fmt.Fprintf(w, "Invalid input distance count")
	}

	limit, err := strconv.ParseInt(arg[0], 10, 32)
	if err != nil {
		fmt.Fprintf(w, "Invalid input distance")
		return
	}

	list := findClosest(&Point{[]float64{latitude, longitude}}, dist, int(limit))

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
