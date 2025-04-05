package main

import (
	"filshr/middleware"
	"filshr/routes"
	"filshr/utils"
	"fmt"
	"net/http"
)

func main() {
	db := utils.ConnectDatabase()

	router := http.NewServeMux()

	router.HandleFunc("POST /upload", routes.UploadHandler(db))
	router.HandleFunc("GET /", routes.GetAsset)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server running on port 8080")
	server.ListenAndServe()
}
