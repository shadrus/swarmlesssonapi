package main

import (
	"encoding/json"
	"net/http"
	"os"
)

var (
	serviceVersion = "0.2"
	token          = "55555dfhigf8-ergush347864t88457gje"
	database       = "db_user:password@tcp(localhost:3306)/my_db"
)

// ServiceInfo is the model for application version info
type ServiceInfo struct {
	Version string
}

func version(w http.ResponseWriter, req *http.Request) {
	info := ServiceInfo{Version: serviceVersion}
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
