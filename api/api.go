package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

type CustomError struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := CustomError{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var RequestErrorHandler = func(w http.ResponseWriter, err error) { writeError(w, err.Error(), http.StatusBadRequest) }
var InternalErrorHandler = func(w http.ResponseWriter) {
	writeError(w, "Internal Server Error", http.StatusInternalServerError)
}
