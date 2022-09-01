package main

import (
	"betting_server/betting_store"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
)

var store = betting_store.NewBettingStore()

func dataHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/getMelbetLink" {
		getMelbetLink(w, req)
	} else if req.URL.Path == "/get1XbetLink" {
		get1xbetLink(w, req)
	} else if req.URL.Path == "/get1winLink" {
		get1winLink(w, req)
	} else if req.URL.Path == "/setReviewValue" {
		setReviewValue(w, req)
	} else {
		http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}
}

func get1winLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1win link %s\n", req.URL.Path)

	winBetLink := store.GetWinLink()
	js, err := json.Marshal(winBetLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func getMelbetLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get Melbet link %s\n", req.URL.Path)

	melBetLink := store.GetMelBetLink()
	js, err := json.Marshal(melBetLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func get1xbetLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1xbetLink %s\n", req.URL.Path)

	betLink := store.Get1XBetLink()
	js, err := json.Marshal(betLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func setReviewValue(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task create at %s\n", req.URL.Path)
	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var rt struct {
		NewReviewValue bool `json:"newValue"`
	}
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.ChangeReviewValue(rt.NewReviewValue)
	js, err := json.Marshal("Success")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func main() {
	port := ":8000"
	osPort, ok := os.LookupEnv("ServerPort")
	if ok {
		port = ":" + osPort
	}
	http.HandleFunc("/", dataHandler) //http://127.0.0.1:8000/get1winLink to get data
	log.Fatal(http.ListenAndServe(port, nil))
}
