package song_services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ExternalAPI struct {
	ReleaseDate string `json:"release_date" db:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func fetchFromLastFM(group, song string) (string, string, error) {
	apiKey := "402e569eec4332ba53a56a0b377e5e13"

	groupEncoded := url.QueryEscape(group)
	songEncoded := url.QueryEscape(song)

	apiURL := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=track.getInfo&api_key=%s&artist=%s&track=%s&format=json",
		apiKey, groupEncoded, songEncoded,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("Last.fm API вернул код ошибки %d", resp.StatusCode)
	}

	var result struct {
		Track struct {
			URL   string `json:"url"`
			Album struct {
				ReleaseDate string `json:"release_date"`
			} `json:"album"`
			Wiki struct {
				Published string `json:"published"`
			} `json:"wiki"`
		} `json:"track"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	link := result.Track.URL

	releaseDate := result.Track.Album.ReleaseDate
	if releaseDate == "" {
		releaseDate = result.Track.Wiki.Published
	}

	return releaseDate, link, nil
}

func fetchLyrics(group, song string) (string, error) {
	groupEncoded := url.QueryEscape(group)
	songEncoded := url.QueryEscape(song)

	apiURL := fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", groupEncoded, songEncoded)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Lyrics.ovh API вернул код ошибки %d", resp.StatusCode)
	}

	var result struct {
		Lyrics string `json:"lyrics"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Lyrics, nil
}

func (s *SongService) fetchSongDetails(group, song string) (*ExternalAPI, error) {
	const op = "song_services.fetchSongDetails"

	releaseDate, link, err := fetchFromLastFM(group, song)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	text, err := fetchLyrics(group, song)
	if err != nil {
		s.l.Warn("Не удалось получить текст песни", "ошибка", err)
		text = ""
	}

	return &ExternalAPI{
		ReleaseDate: releaseDate,
		Text:        text,
		Link:        link,
	}, nil
}
