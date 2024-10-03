package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	u, err := h.auth.GetSessionUser(r)
	if err != nil {
		fmt.Println("There are no authenticated user")
		return
	}
	WriteJSON(w, http.StatusOK, u)
}
