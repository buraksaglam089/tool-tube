package playlist

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	spotifyTokenURL = "https://accounts.spotify.com/api/token"
	spotifyAPIURL   = "https://api.spotify.com/v1"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type PlaylistResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tracks      struct {
		Items []struct {
			Track struct {
				Name    string `json:"name"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
}

type PlaylistInfo struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Songs       []string `json:"songs"`
}

// getToken retrieves an access token from Spotify
func getToken(ctx context.Context) (string, error) {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		return "", fmt.Errorf("missing Spotify credentials")
	}

	auth := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, "POST", spotifyTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get token: %s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read token response: %w", err)
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal token response: %w", err)
	}

	return tokenResp.AccessToken, nil
}

// getPlaylist retrieves playlist information from Spotify
func getPlaylist(ctx context.Context, token, playlistID string) (*PlaylistResponse, error) {
	url := fmt.Sprintf("%s/playlists/%s", spotifyAPIURL, playlistID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create playlist request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get playlist: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get playlist: %s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read playlist response: %w", err)
	}

	var playlist PlaylistResponse
	err = json.Unmarshal(body, &playlist)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal playlist response: %w", err)
	}

	return &playlist, nil
}

// GetPlaylistSong retrieves playlist information including songs, title, and description
func GetPlaylistSong(playlistID string) (*PlaylistInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token, err := getToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting token: %v", err)
	}

	playlist, err := getPlaylist(ctx, token, playlistID)
	if err != nil {
		return nil, fmt.Errorf("error getting playlist: %v", err)
	}

	log.Println("Playlist response:", playlist)

	if len(playlist.Tracks.Items) == 0 {
		return &PlaylistInfo{
			Title:       playlist.Name,
			Description: playlist.Description,
			Songs:       []string{},
		}, nil
	}

	var songs []string
	for _, item := range playlist.Tracks.Items {
		track := item.Track
		artists := make([]string, len(track.Artists))
		for j, artist := range track.Artists {
			artists[j] = artist.Name
		}
		songInfo := fmt.Sprintf("%s by %s", track.Name, strings.Join(artists, ", "))
		songs = append(songs, songInfo)
	}

	playlistInfo := &PlaylistInfo{
		Title:       playlist.Name,
		Description: playlist.Description,
		Songs:       songs,
	}

	return playlistInfo, nil
}
