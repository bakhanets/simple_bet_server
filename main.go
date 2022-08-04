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

type taskServer struct {
	store *betting_store.BettingStore
}

func NewTaskServer() *taskServer {
	store := betting_store.New()
	return &taskServer{store: store}
}

func (ts *taskServer) dataHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/getMelbetLink" {
		ts.getMelbetLink(w, req)
	} else if req.URL.Path == "/get1XbetLink" {
		ts.get1xbetLink(w, req)
	} else if req.URL.Path == "/get1winLink" {
		ts.get1winLink(w, req)
	} else if req.URL.Path == "/getReviewValue" {
		ts.getReviewValue(w, req)
	} else if req.URL.Path == "/setReviewValue" {
		ts.setReviewValue(w, req)
	} else {
		http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (ts *taskServer) getReviewValue(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1xbetLink %s\n", req.URL.Path)

	reviewValue := ts.store.GetReviewValue()
	js, err := json.Marshal(reviewValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) get1winLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1win link %s\n", req.URL.Path)

	winBetLink := ts.store.GetWinLink()
	js, err := json.Marshal(winBetLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) getMelbetLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get Melbet link %s\n", req.URL.Path)

	melBetLink := ts.store.GetMelbetLink()
	js, err := json.Marshal(melBetLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) get1xbetLink(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get 1xbetLink %s\n", req.URL.Path)

	betLink := ts.store.Get1XbetLink()
	js, err := json.Marshal(betLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) setReviewValue(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task create at %s\n", req.URL.Path)

	type RequestBody struct {
		NewReviewValue bool `json:"newValue"`
	}

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
	var rt RequestBody
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts.store.ChangeReviewValue(rt.NewReviewValue)
	js, err := json.Marshal("Success")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()
	mux.HandleFunc("/", server.dataHandler)

	log.Fatal(http.ListenAndServe(":8000"+os.Getenv("SERVERPORT"), mux))
}
