package controller

import (
	"encoding/json"
	"fmt"
	"imports/dependencies/app/model"
	"net/http"
	"strings"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		w.Header().Set("Content-Type", "application/json")
		eneableCors(&w)

		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {

			limite := "5"
			desde := "0"

			limiteGet, ok := r.URL.Query()["limite"]

			if ok {
				limite = limiteGet[0]
			}

			desdeGet, ok := r.URL.Query()["desde"]

			if ok {
				desde = desdeGet[0]
			}

			videos, err := model.GetAllVideos(limite, desde)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				json.NewEncoder(w).Encode(videos)
			}
		}
	}
}

func GetAllVideosPadre(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		w.Header().Set("Content-Type", "application/json")
		eneableCors(&w)

		authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

		tokenDetails, err := model.ValidateToken(authToken)

		if err != nil {
			fmt.Print("Invalid Token")
		}

		userPadre := tokenDetails.USUARIO.PADRE
		if userPadre == "" {
			fmt.Print("The user does not have a father")
		}

		limite := "5"
		desde := "0"

		limiteGet, ok := r.URL.Query()["limite"]

		if ok {
			limite = limiteGet[0]
		}

		desdeGet, ok := r.URL.Query()["desde"]

		if ok {
			desde = desdeGet[0]
		}

		videos, err := model.GetAllVideosPadre(userPadre, limite, desde)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			json.NewEncoder(w).Encode(videos)
		}
	}
}
