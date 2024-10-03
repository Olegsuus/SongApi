package models

//type AddSongDTO struct {
//	Group string `json:"group"`
//	Song  string `json:"song"`
//}
//
//type UpdateSongDTO struct {
//	Group       string `json:"group,omitempty"`
//	Song        string `json:"song,omitempty"`
//	ReleaseDate string `json:"release_date,omitempty"`
//	Text        string `json:"text,omitempty"`
//	Link        string `json:"link,omitempty"`
//}

type UpdateSongDTO struct {
	Group       string `json:"group,omitempty" validate:"omitempty,max=255,not_russian"`
	Song        string `json:"song,omitempty" validate:"omitempty,max=255,not_russian"`
	ReleaseDate string `json:"release_date,omitempty" validate:"omitempty,max=50"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty" validate:"omitempty,max=255"`
}

type AddSongDTO struct {
	Group string `json:"group" validate:"required,max=255,not_russian"`
	Song  string `json:"song" validate:"required,max=255,not_russian"`
}
