package handlers

import (
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	err := WriteJSON(w, http.StatusOK, "burak")
	return err
}
