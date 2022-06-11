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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			users, err := model.GetAllUsers()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				json.NewEncoder(w).Encode(users)
			}
		}
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			param := mux.Vars(r)["uid"]
			uid, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			user, err := model.GetUser(uid)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			json.NewEncoder(w).Encode(user)
		}
	}
}

/*
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
*/

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			decoder := json.NewDecoder(r.Body)
			var user model.User
			err := decoder.Decode(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			updatedUser, err := model.UpdateUser(user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				w.WriteHeader(http.StatusOK)

				fmt.Print("Se supone que se ha actualizado")

				enc := json.NewEncoder(w)
				enc.SetIndent("", "  ")
				enc.Encode(updatedUser)

			}
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			param := mux.Vars(r)
			idStr := param["id"]
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			err = model.DeleteUser(id)
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

func RegistrationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			dec := json.NewDecoder(r.Body)
			var user model.User
			err := dec.Decode(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			if user.NOMBRE == "" || user.CORREO == "" || user.PASSWORD == "" {
				fmt.Fprintf(w, "Please enter a valid username and password.\r\n")
			} else {
				response, err := model.CreateUser(user.NOMBRE, user.CORREO, user.PASSWORD)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				} else {
					fmt.Fprintf(w, response)
				}
			}
		}
	}
}

func AuthenticationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		eneableCors(&w)

		fmt.Print(w.Header())

		dec := json.NewDecoder(r.Body)
		var user model.User
		err := dec.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if user.NOMBRE == "" || user.PASSWORD == "" {
			fmt.Fprintf(w, "Please enter a valid username and password.\r\n")
		} else {
			tokenDetails, err := model.GenerateToken(user.NOMBRE, user.PASSWORD)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				enc := json.NewEncoder(w)
				enc.SetIndent("", "  ")
				enc.Encode(tokenDetails)
			}
		}
	}
}

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			fmt.Println(w.Header())

			authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

			fmt.Println("* El token es " + authToken)

			userDetails, err := model.ValidateToken(authToken)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				enc := json.NewEncoder(w)
				enc.Encode(userDetails)
			}
		}
	}
}
