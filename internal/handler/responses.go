package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func errorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.Header().Set("X=Content-Type-Options", "nosniff")

	w.WriteHeader(statusCode)
	js := fmt.Sprintf(`{"error":"%s"}`, message)
	fmt.Fprint(w, js)
}

// jsonResponse function response with the json representing the interface v.
func jsonResponse(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}
