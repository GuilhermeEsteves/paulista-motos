package login

import (
	"encoding/json"
	"net/http"
)

type responseLogin struct {
	Token string `json:"token"`
}

// AuthenticationHandler método valida o usuario e senha e devolve um token
func AuthenticationHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(responseLogin{
		Token: "123",
	})
}
