package song_services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ExternalAPI struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

//func (s *SongService) fetchSongDetails(group, song string) (*ExternalAPI, error) {
//	const op = "song_services.fetchSongDetails"
//
//	groupEncoded := url.QueryEscape(group)
//	songEncoded := url.QueryEscape(song)
//
//	// Формирование URL согласно Swagger
//	apiURL := fmt.Sprintf("http://localhost:8888/info?group=%s&song=%s", groupEncoded, songEncoded)
//
//	// Создание запроса
//	req, err := http.NewRequest("GET", apiURL, nil)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("%s: external API returned status code %d", op, resp.StatusCode)
//	}
//
//	// Парсинг ответа
//	var apiResponse ExternalAPI
//	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
//		return nil, fmt.Errorf("%s: failed to decode API response: %w", op, err)
//	}
//
//	return &apiResponse, nil
//}

func fetchFromLastFM(group, song string) (string, string, error) {
	apiKey := "402e569eec4332ba53a56a0b377e5e13" // Замените на ваш API ключ Last.fm

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
				ReleaseDate string `json:"releasedate"`
			} `json:"album"`
		} `json:"track"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	releaseDate := result.Track.Album.ReleaseDate
	link := result.Track.URL

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

// Функция для объединения данных из Last.fm и Lyrics.ovh
func (s *SongService) fetchSongDetails(group, song string) (*ExternalAPI, error) {
	const op = "song_services.fetchSongDetails"

	// Получаем данные из Last.fm
	releaseDate, link, err := fetchFromLastFM(group, song)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Получаем текст песни из Lyrics.ovh
	text, err := fetchLyrics(group, song)
	if err != nil {
		s.l.Warn("Не удалось получить текст песни", "ошибка", err)
		// Если текст не найден, оставляем пустую строку
		text = ""
	}

	// Возвращаем обогащенные данные
	return &ExternalAPI{
		ReleaseDate: releaseDate,
		Text:        text,
		Link:        link,
	}, nil
}
