package models

type AddSongDTO struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type UpdateSongDTO struct {
	Group       string `json:"group,omitempty"`
	Song        string `json:"song,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}
