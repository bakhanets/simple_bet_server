package main

import (
	"betting_server/links_storage"
	"encoding/json"
	"github.com/oschwald/maxminddb-golang"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"strings"
)

var storage links_storage.Storage

const geoLiteCountryDBPath = "/geoip/GeoLite2-Country.mmdb"

// legacy block start
func get1Link(w http.ResponseWriter, req *http.Request) {
	writeResponse(w, req, "v1")
}

func get2Link(w http.ResponseWriter, req *http.Request) {
	writeResponse(w, req, "v2")
}

func get3Link(w http.ResponseWriter, req *http.Request) {
	writeResponse(w, req, "v3")
}

func get4Link(w http.ResponseWriter, req *http.Request) {
	writeResponse(w, req, "v4")
} // legacy block end

func handleFunc(w http.ResponseWriter, req *http.Request) {
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
	log.Printf("handling get link %s\nFrom %s", req.URL.Path, req.RemoteAddr)
	value, ok := storage.GetValueByKeyForCountry(key, getIsoCountryNameFromIp(req.RemoteAddr))
	if !ok {
		http.Error(w, "No value for this key", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(value))
}

func getIsoCountryNameFromIp(remoteAddr string) (isoCode string) {
	defer func() {
		log.Println("got", isoCode, "for", remoteAddr)
	}()
	ipAndPort := strings.Split(remoteAddr, ":")
	if len(ipAndPort) != 2 {
		log.Println("error getting ip from \"" + remoteAddr + "\"")
		return ""
	}
	ip := net.ParseIP(ipAndPort[0])

	db, err := maxminddb.Open(geoLiteCountryDBPath)
	if err != nil {
		log.Println("error opening geoip db:", err)
		return ""
	}
	defer func(db *maxminddb.Reader) {
		err := db.Close()
		if err != nil {
			log.Println("error closing GeoIp db:", err)
		}
	}(db)

	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	err = db.Lookup(ip, &record)
	if err != nil {
		log.Println("error searching for ip in geoip db:", err)
		return ""
	}
	return record.Country.ISOCode
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
	http.HandleFunc("/getValue", handleFunc)
	http.HandleFunc("/setReviewValue", setReviewValue)
	log.Fatal(http.ListenAndServe(port, nil))
}
