package controller

import (
	"fmt"
	controller "imports/dependencies/app/controller/web"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/usuarios", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/usuario/{uid}", controller.GetUser).Methods("GET")

	//	router.HandleFunc("/api/usuarios", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/usuarios", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/usuarios/{id}", controller.DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/registrations", controller.RegistrationsHandler).Methods("POST")
	router.HandleFunc("/api/authentications", controller.AuthenticationsHandler).Methods("POST")
	router.HandleFunc("/api/test", controller.TestResourceHandler)
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
