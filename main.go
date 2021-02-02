package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// ServiceInfo is the model for application version info
type ServiceInfo struct {
	Version string
}

func version(w http.ResponseWriter, req *http.Request) {
	info := ServiceInfo{Version: "0.2"}
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func exit(w http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func crash(w http.ResponseWriter, req *http.Request) {
	os.Exit(1)
}

func main() {
	http.HandleFunc("/version", version)
	http.HandleFunc("/exit", exit)
	http.HandleFunc("/crash", crash)
	http.ListenAndServe(":80", nil)
}
