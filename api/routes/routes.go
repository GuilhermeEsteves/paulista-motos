package routes

import (
	"net/http"

	"github.com/GuilhermeEsteves/paulista-motos/api/login"
	"github.com/GuilhermeEsteves/paulista-motos/api/ping"
	"github.com/gorilla/mux"
)

// ConfigureRoutes método que configura as rotas da api
func ConfigureRoutes() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ping", ping.GetHandler).Methods("GET")
	api.HandleFunc("/login", login.AuthenticationHandler).Methods("POST")

	frontEnd := router.PathPrefix("/")
	frontEnd.Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("./front/"))))

	return router
}
