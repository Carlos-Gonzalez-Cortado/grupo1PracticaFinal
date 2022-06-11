package controller

import (
	"encoding/json"
	"fmt"
	"imports/dependencies/app/model"
	"net/http"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		eneableCors(&w)
	} else {
		if !TokenCheck(w, r) {
			fmt.Print("Invalid Token")
		} else {
			w.Header().Set("Content-Type", "application/json")
			eneableCors(&w)

			videos, err := model.GetAllVideos()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				json.NewEncoder(w).Encode(videos)
			}
		}
	}
}
