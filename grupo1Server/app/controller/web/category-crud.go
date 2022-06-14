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

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
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

			categories, err := model.GetAllCategories(limite, desde)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				json.NewEncoder(w).Encode(categories)
			}
		}
	}
}

func GetAllCategoriesPadre(w http.ResponseWriter, r *http.Request) {
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

		categories, err := model.GetAllCategoriesPadre(userPadre, limite, desde)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			json.NewEncoder(w).Encode(categories)
		}
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
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

			err = model.DeleteCategory(id)
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

func CreateCategory(w http.ResponseWriter, r *http.Request) {
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
			var categoria model.Tipos
			err := dec.Decode(&categoria)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			if categoria.NOMBRE == "" {
				fmt.Fprintf(w, "Please enter a valid category name.\r\n")
			} else {
				response, err := model.CreateCategory(categoria.NOMBRE, userId)
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

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
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
			var categoria model.Tipos
			errDec := decoder.Decode(&categoria)

			if errDec != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errDec.Error()))
				return
			}

			/// ->
			categoria.ID = id
			/// <-

			updatedCategory, err := model.UpdateCategory(categoria)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.WriteHeader(http.StatusOK)

				fmt.Print("Se supone que se ha actualizado")

				enc := json.NewEncoder(w)
				enc.SetIndent("", "  ")
				enc.Encode(updatedCategory)
			}
		}
	}
}
