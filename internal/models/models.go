package models

import "time"

type Song struct {
	ID          int       `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate string    `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AddSongDTO struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type RemoveDTO struct {
	ID int `json:"id" validate:"required"`
}

type UpdateSongDTO struct {
	Group       string `json:"group" validate:"required"`
	Song        string `json:"song" validate:"required"`
	ReleaseDate string `json:"release_date" validate:"required"`
	Text        string `json:"text"`
	Link        string `json:"link" validate:"required"`
}

type GetManySong struct {
	Group       string `query:"group"`
	Song        string `query:"song"`
	ReleaseDate string `query:"release_date"`
	Text        string `query:"text"`
	Link        string `query:"link"`
	Size        int    `query:"size"`
	Page        int    `query:"page"`
	SortBy      string `query:"sort_by"`    // Новое поле
	SortOrder   string `query:"sort_order"` // Новое поле ("asc" или "desc")
}

type GetManySongs struct {
	Songs []*Song `json:"songs"`
	Page  int     `json:"page"`
	Size  int     `json:"size"`
}

type GetSongText struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

type SongText struct {
	ID     int      `json:"id"`
	Group  string   `json:"group"`
	Song   string   `json:"song"`
	Lyrics []string `json:"lyrics"`
	Page   int      `json:"page"`
	Size   int      `json:"size"`
	Total  int      `json:"total"`
}
