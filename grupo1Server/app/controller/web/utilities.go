package controller

import (
	"fmt"
	"imports/dependencies/app/model"
	"net/http"
	"strings"
)

func eneableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func TokenCheck(w http.ResponseWriter, r *http.Request) bool {
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	userDetails, err := model.ValidateToken(authToken)

	fmt.Println("* Datos del usuario: " + userDetails.USUARIO.ROL)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return false
	} else if userDetails.USUARIO.ROL != "ADMIN_ROL" {
		return false
	} else {
		return true
	}
}
