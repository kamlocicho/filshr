package routes

import (
	"database/sql"
	"filshr/services"
	"filshr/utils"
	"fmt"
	"net/http"
)

type Response struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func UploadHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			// Accept only POST
			http.Error(w, "This method is not allowed.", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("Uploading a new file")

		fileName, err := services.FileUpload(r, "./storage/")
		if err != nil {
			http.Error(w, fmt.Sprintf("There was an issue uploading this file: %s", err), http.StatusBadRequest)
			return
		}

		utils.RespondJson(w, Response{
			Name:   fileName,
			Status: http.StatusOK,
		})
	}
}
