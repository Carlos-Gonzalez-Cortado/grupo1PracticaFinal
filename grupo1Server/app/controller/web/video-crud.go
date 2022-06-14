package controller

import (
	"encoding/json"
	"fmt"
	"imports/dependencies/app/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			param := mux.Vars(r)["id"]
			id, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			err = model.DeleteVideo(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

func CreateVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

			tokenDetails, errUser := model.ValidateToken(authToken)

			if errUser != nil {
				fmt.Print("Invalid Token")
			}

			userId := tokenDetails.USUARIO.UID
			if userId < 0 {
				fmt.Print("The user is no valid")
			}

			dec := json.NewDecoder(r.Body)
			var video model.SimpleVideo
			err := dec.Decode(&video)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			if video.NOMBRE == "" || video.URL == "" || video.CATEGORIA < 0 {
				fmt.Fprintf(w, "Please enter a valid video.\r\n")
			} else {
				response, err := model.CreateVideo(video.NOMBRE, video.URL, userId, video.CATEGORIA)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				} else {
					enc := json.NewEncoder(w)
					enc.SetIndent("", "  ")
					enc.Encode(response)
				}
			}
		}
	}
}

func UpdateVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			param := mux.Vars(r)["id"]
			id, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			/// <-

			decoder := json.NewDecoder(r.Body)
			var video model.SimpleVideo
			errDec := decoder.Decode(&video)

			if errDec != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errDec.Error()))
				return
			}

			/// ->
			video.ID = id
			/// <-

			updatedVideo, err := model.UpdateVideo(video)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.WriteHeader(http.StatusOK)

				fmt.Print("Se supone que se ha actualizado")

				enc := json.NewEncoder(w)
				enc.SetIndent("", "  ")
				enc.Encode(updatedVideo)

			}
		}
	}
}
