package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", healthcheck)

	http.ListenAndServe(":8080", nil)

}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
