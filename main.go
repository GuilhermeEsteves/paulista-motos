package main

import (
	"log"
	"net/http"

	"github.com/GuilhermeEsteves/paulista-motos/api/routes"
)

func main() {
	router := routes.ConfigureRoutes()

	server := &http.Server{
		Handler: router,
		Addr:    ":5000",
	}

	log.Println("API starting at port 5000 :)")
	log.Fatal(server.ListenAndServe())

}
