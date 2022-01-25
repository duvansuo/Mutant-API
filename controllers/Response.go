package controllers

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, http_code int, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
