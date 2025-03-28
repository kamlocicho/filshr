package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondJson(w http.ResponseWriter, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse object to json: %s", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
