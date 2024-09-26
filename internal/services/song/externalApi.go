package song_services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExternalAPI struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"patronymic"`
}

func (s *SongService) fetchSongDetails(group, song string) (*ExternalAPI, error) {
	const op = "song_services.fetchSongDetails"

	apiURL := fmt.Sprintf("http://external-api.com/info?group=%s&song=%s", group, song)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: external API returned status code %d", op, resp.StatusCode)
	}

	var apiResponse ExternalAPI
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("%s: failed to decode API response: %w", op, err)
	}

	s.l.Info("Successfully fetched song details from external API", "group", group, "song", song)

	return &apiResponse, nil
}
