package main

import (
	"betting_server/betting_store"
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"os"
)

var store = betting_store.NewBettingStore()

func get1winLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1win link %s\n", req.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(store.GetWinLink()))
}

func getMelbetLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get Melbet link %s\n", req.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(store.GetMelBetLink()))
}

func get1xbetLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1xbetLink %s\n", req.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(store.Get1XBetLink()))
}

func setReviewValue(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task create at %s\n", req.URL.Path)
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
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("\"Success\""))
}

func main() {
	port := ":8000"
	osPort, ok := os.LookupEnv("ServerPort")
	if ok {
		port = ":" + osPort
	}
	http.HandleFunc("/v1/getPredictionsList", get1xbetLink)
	http.HandleFunc("/v2/getPredictionsList", get1winLink)
	http.HandleFunc("/v3/getPredictionsList", getMelbetLink)
	http.HandleFunc("/setReviewValue", setReviewValue)
	log.Fatal(http.ListenAndServe(port, nil))
}
