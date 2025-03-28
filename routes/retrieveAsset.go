package routes

import (
	"fmt"
	"net/http"
)

func GetAsset(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Responding with an image")
}
