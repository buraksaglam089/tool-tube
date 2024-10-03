package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/buraksaglam089/tool-tube/services/playlist"
	"github.com/buraksaglam089/tool-tube/types"
)

type RequestData struct {
	ID string `json:"playlistId"`
}

func (h *Handler) ConvertPlaylist(w http.ResponseWriter, r *http.Request) {
	u, err := h.auth.GetSessionUser(r)
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var data RequestData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Cannot decode request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if data.ID == "" {
		http.Error(w, "Playlist ID is required", http.StatusBadRequest)
		return
	}

	//
	spotifyPlaylist, err := playlist.GetPlaylistSong(data.ID)
	if err != nil {
		log.Println("Failed to get Spotify playlist:", err)
		http.Error(w, "Failed to get Spotify playlist", http.StatusInternalServerError)
		return
	}
	var user types.User
	if err := h.DB.Where("email = ?", u.Email).First(&user).Error; err != nil {
		log.Println("Failed to find user:", err)
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		return
	}

	accessToken := user.AccessToken
	if accessToken == "" {
		http.Error(w, "YouTube access token is missing", http.StatusUnauthorized)
		return
	}

	err = playlist.ConvertSpotifyToYouTubePlaylist(accessToken, *spotifyPlaylist, "public")
	if err != nil {
		log.Println("Failed to convert playlist:", err)
		http.Error(w, "Failed to convert playlist", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Playlist converted successfully"}
	WriteJSON(w, http.StatusOK, response)
}
