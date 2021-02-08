package main

import (
	"encoding/json"
	"net/http"
	"os"

	secrets "github.com/ijustfool/docker-secrets"
)

// ServiceInfo is the model for application version info
type ServiceInfo struct {
	Version string
}

// Ручка токен получает токен из секрета Docker, сверяет с токеном из запроса.
func token(w http.ResponseWriter, req *http.Request) {
	// /run/secrets/<secret_name> for secrets and /<config-name> for configs
	reqToken, ok := req.URL.Query()["token"]

	if !ok || len(reqToken[0]) < 1 {
		http.Error(w, "No token in your reguest", http.StatusBadRequest)
		return
	}
	dockerSecrets, _ := secrets.NewDockerSecrets("")
	token, _ := dockerSecrets.Get("token")
	if token != reqToken[0] {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}

// Ручка version получает номер версии из переменных окружения.
func version(w http.ResponseWriter, req *http.Request) {
	info := ServiceInfo{Version: os.Getenv("vesrion")}
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
	http.HandleFunc("/token", token)
	http.ListenAndServe(":80", nil)
}
