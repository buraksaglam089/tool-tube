package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/buraksaglam089/tool-tube/types"
)

func (h *Handler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := h.DB.Create(&user).Error; err != nil {
		http.Error(w, "Can not create", http.StatusBadRequest)
		return
	}

	WriteJSON(w, http.StatusCreated, user)
}

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	u, err := h.auth.GetSessionUser(r)
	if err != nil {
		fmt.Println("There are no authenticated user")
		return
	}
	WriteJSON(w, http.StatusOK, u)
}
