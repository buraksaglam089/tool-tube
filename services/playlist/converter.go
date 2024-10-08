package playlist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CreatePlaylistRequest struct {
	Snippet struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"snippet"`
	Status struct {
		PrivacyStatus string `json:"privacyStatus"`
	} `json:"status"`
}
type CreatePlaylistResponse struct {
	ID string `json:"id"`
}

func createYoutubePlaylist(accessToken, title, description, privacyStatus string) (string, error) {
	playlistRequest := CreatePlaylistRequest{}
	playlistRequest.Snippet.Title = title
	playlistRequest.Snippet.Description = description
	playlistRequest.Status.PrivacyStatus = privacyStatus

	jsonData, err := json.Marshal(playlistRequest)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	apiURL := "https://www.googleapis.com/youtube/v3/playlists?part=snippet%2Cstatus"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	log.Println(resp, "resposne of xdddddd   ")
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return "", fmt.Errorf("failed to create playlist, status code: %d, response: %s", resp.StatusCode, (resp.Body))
	}

	var createPlaylistResponse CreatePlaylistResponse
	err = json.NewDecoder(resp.Body).Decode(&createPlaylistResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	return createPlaylistResponse.ID, nil
}

func addSongToPlaylist(accessToken, playlistID, videoID string) error {

	addRequest := struct {
		Snippet struct {
			PlaylistID string `json:"playlistId"`
			ResourceID struct {
				Kind    string `json:"kind"`
				VideoID string `json:"videoId"`
			} `json:"resourceId"`
		} `json:"snippet"`
	}{}

	addRequest.Snippet.PlaylistID = playlistID
	addRequest.Snippet.ResourceID.Kind = "youtube#video"
	addRequest.Snippet.ResourceID.VideoID = videoID

	jsonData, err := json.Marshal(addRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	apiURL := "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to add song to playlist, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

func searchYouTube(accessToken, query string) (string, error) {
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&maxResults=1&q=%s", url.QueryEscape(query))
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to search YouTube, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var searchResponse struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
		} `json:"items"`
	}

	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	if len(searchResponse.Items) == 0 {
		return "", fmt.Errorf("no results found for query: %s", query)
	}

	videoID := searchResponse.Items[0].ID.VideoID
	return videoID, nil
}
func ConvertSpotifyToYouTubePlaylist(accessToken string, spotifyPlaylist PlaylistInfo, privacyStatus string) error {

	fmt.Printf("Processing Spotify playlist: %s\n", spotifyPlaylist.Title)

	youtubePlaylistID, err := createYoutubePlaylist(accessToken, spotifyPlaylist.Title, spotifyPlaylist.Description, privacyStatus)
	if err != nil {
		return fmt.Errorf("failed to create YouTube playlist: %v", err)
	}

	fmt.Printf("Created YouTube playlist with ID: %s\n", youtubePlaylistID)

	for _, song := range spotifyPlaylist.Songs {

		query := song

		videoID, err := searchYouTube(accessToken, query)
		if err != nil {
			fmt.Printf("Failed to find YouTube video for '%s': %v\n", query, err)
			continue
		}

		err = addSongToPlaylist(accessToken, youtubePlaylistID, videoID)
		if err != nil {
			fmt.Printf("Failed to add video '%s' to playlist: %v\n", videoID, err)
			continue
		}

		fmt.Printf("Added '%s' to YouTube playlist\n", query)
	}

	fmt.Println("Conversion complete!")
	return nil
}
