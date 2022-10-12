package main

import (
	"betting_server/links_storage"
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"os"
)

var storage links_storage.Storage

// legacy block start
func get1Link(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get v1 link")
	writeResponse(w, req, "v1")
}

func get2Link(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get v2 link")
	writeResponse(w, req, "v2")
}

func get3Link(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get v3 link")
	writeResponse(w, req, "v3")
}

func get4Link(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get v4 link")
	writeResponse(w, req, "v4")
} // legacy block end

func handleFunc(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get link %s\n", req.URL.Path)
	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !req.Form.Has("key") {
		http.Error(w, "Error: value for \"key\" not found", http.StatusBadRequest)
		return
	}
	writeResponse(w, req, req.Form.Get("key"))
}

func writeResponse(w http.ResponseWriter, req *http.Request, key string) {
	log.Println(req.RemoteAddr)
	value, ok := storage.GetValueByKey(key)
	if !ok {
		http.Error(w, "No value for this key", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(value))
}

func setReviewValue(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task create at %s\n", req.URL.Path)
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
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

	storage.SetReviewValue(rt.NewReviewValue)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("\"Success\""))
}

func main() {
	port := ":8000"
	osPort, ok := os.LookupEnv("ServerPort")
	if ok {
		port = ":" + osPort
	}
	err := storage.LoadValues()
	if err != nil {
		log.Fatalln(err)
	}
	storage.SetReviewValue(true)
	// legacy block start
	http.HandleFunc("/v1/getPredictionsList", get1Link)
	http.HandleFunc("/v2/getPredictionsList", get2Link)
	http.HandleFunc("/v3/getPredictionsList", get3Link)
	http.HandleFunc("/v4/getPredictionsList", get4Link)
	// legacy block end
	http.HandleFunc("/getPredictionsList", handleFunc)
	http.HandleFunc("/setReviewValue", setReviewValue)
	log.Fatal(http.ListenAndServe(port, nil))
}
