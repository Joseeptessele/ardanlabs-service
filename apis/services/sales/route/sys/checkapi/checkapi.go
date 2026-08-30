package checkapi

import (
	"encoding/json"
	"net/http"
)

func liveness(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}

func readiness(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
