package ping

import (
	"encoding/json"
	"net/http"
)

// GetHandler método de health check
func GetHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("pong")
}
