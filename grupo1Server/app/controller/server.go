package controller

import (
	"fmt"
	controller "imports/dependencies/app/controller/web"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initUserHandlers() {
	router.HandleFunc("/api/usuarios", controller.GetAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/usuario/{uid}", controller.GetUser).Methods("GET", "OPTIONS")

	//	router.HandleFunc("/api/usuarios", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/modificar", controller.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/usuarios/{id}", controller.DeleteUser).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/registrar", controller.RegistrationsHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/autentificar", controller.AuthenticationsHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/validar", controller.ValidateTokenHandler).Methods("POST", "OPTIONS")
}

func initVideoHandlers() {
	router.HandleFunc("/api/videos", controller.GetAllVideos).Methods("GET", "OPTIONS")
}

func Start() {
	router = mux.NewRouter()

	initUserHandlers()
	initVideoHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
