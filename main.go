package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type ServiceInfo struct {
	Version string
}

func hello(w http.ResponseWriter, req *http.Request) {
	info := ServiceInfo{Version: "0.1"}
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func crash(w http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func main() {
	http.HandleFunc("/version", hello)
	http.HandleFunc("/crash", crash)
	http.ListenAndServe(":80", nil)
}
