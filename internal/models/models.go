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
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
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
	Size        int    `query:"size,default=5"`
	Page        int    `query:"page,default=1"`
}

type GetManySongs struct {
	Songs []*Song `json:"songs"`
	Page  int     `json:"page"`
	Size  int     `json:"size"`
}

type GetSongText struct {
	Page int `query:"page,default=1"`
	Size int `query:"size,default=1"`
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
