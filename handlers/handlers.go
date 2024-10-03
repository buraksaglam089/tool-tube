package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/buraksaglam089/tool-tube/services/auth"
	"github.com/buraksaglam089/tool-tube/types"
	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

type Handler struct {
	DB   *gorm.DB
	auth *auth.AuthService
}

func NewHandler(db *gorm.DB, auth *auth.AuthService) *Handler {
	return &Handler{DB: db, auth: auth}
}

func (h *Handler) HandleFoo(w http.ResponseWriter, r *http.Request) {
	var users []types.User
	if err := h.DB.Find(&users).Error; err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusAccepted, users)
}

func (h *Handler) HandleProvideLogin(w http.ResponseWriter, r *http.Request) {
	provider := "google"
	log.Printf("Provider: %s\n", provider)

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println(gothUser)
	} else {
		log.Printf("Error in CompleteUserAuth: %v\n", err)
		gothic.BeginAuthHandler(w, r)
	}
}

// func (h *Handler) HandleAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
// 	provider := "google"
// 	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
// 	user, err := gothic.CompleteUserAuth(w, r)

// 	if err != nil {
// 		log.Printf("Error completing auth: %v\n", err)
// 		http.Error(w, "Error completing auth", http.StatusInternalServerError)
// 		return
// 	}

// 	// Extract user information from the authentication result
// 	googleID := user.UserID

// 	// Check if the user already exists in the database
// 	var existingUser types.User
// 	if err := h.DB.Where("google_id = ?", googleID).First(&existingUser).Error; err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			// An error occurred while querying the database
// 			log.Printf("Database error: %v\n", err)
// 			http.Error(w, "Database error", http.StatusInternalServerError)
// 			return
// 		}

// 		// User does not exist, so create a new user
// 		newUser := &types.User{
// 			GoogleID:     &user.UserID,
// 			FirstName:    user.FirstName,
// 			LastName:     user.LastName,
// 			Email:        user.Email,
// 			AccessToken:  user.AccessToken,
// 			RefreshToken: user.RefreshToken,
// 		}

// 		if err := h.DB.Create(newUser).Error; err != nil {
// 			log.Printf("Error creating user: %v\n", err)
// 			http.Error(w, "Error creating user", http.StatusInternalServerError)
// 			return
// 		}

// 		existingUser = *newUser
// 	} else {
// 		existingUser.FirstName = user.FirstName
// 		existingUser.LastName = user.LastName
// 		existingUser.Email = user.Email
// 		existingUser.AccessToken = user.AccessToken
// 		existingUser.RefreshToken = user.RefreshToken

// 		if err := h.DB.Save(&existingUser).Error; err != nil {
// 			log.Printf("Error updating user: %v\n", err)
// 			http.Error(w, "Error updating user", http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	// Store the user session
// 	err = h.auth.StoreUserSession(w, r, user)
// 	if err != nil {
// 		log.Printf("Error storing user session: %v\n", err)
// 		http.Error(w, "Error storing user session", http.StatusInternalServerError)
// 		return
// 	}

//		// Redirect or respond as needed
//		log.Printf("User authenticated: %v\n", existingUser)
//		http.Redirect(w, r, "http://localhost:5173/", http.StatusSeeOther)
//	}
func (h *Handler) HandleAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := "google"
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("Error completing auth: %v\n", err)
		http.Error(w, fmt.Sprintf("Error completing auth: %v", err), http.StatusInternalServerError)
		return
	}

	googleID := user.UserID

	var existingUser types.User
	if err := h.DB.Where("google_id = ?", googleID).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Database error: %v\n", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		newUser := &types.User{
			GoogleID:     &user.UserID,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			AccessToken:  user.AccessToken,
			RefreshToken: user.RefreshToken,
		}

		if err := h.DB.Create(newUser).Error; err != nil {
			log.Printf("Error creating user: %v\n", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		existingUser = *newUser
	} else {

		existingUser.FirstName = user.FirstName
		existingUser.LastName = user.LastName
		existingUser.Email = user.Email
		existingUser.AccessToken = user.AccessToken
		existingUser.RefreshToken = user.RefreshToken

		if err := h.DB.Save(&existingUser).Error; err != nil {
			log.Printf("Error updating user: %v\n", err)
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			return
		}
	}

	err = h.auth.StoreUserSession(w, r, user)
	if err != nil {
		log.Printf("Error storing user session: %v\n", err)
		http.Error(w, "Error storing user session", http.StatusInternalServerError)
		return
	}

	log.Printf("User authenticated: %v\n", existingUser)
	http.Redirect(w, r, "http://localhost:5173/", http.StatusSeeOther)
}
