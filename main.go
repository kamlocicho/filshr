package main

import (
	"filshr/routes"
	"filshr/utils"
	"log"
	"net/http"
)

func main() {
	db := utils.ConnectDatabase()
	http.HandleFunc("/upload", routes.UploadHandler(db))
	http.HandleFunc("/", routes.GetAsset)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting up the server: ", err)
	}
}
