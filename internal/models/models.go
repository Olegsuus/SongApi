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

type GetManySong struct {
	Group       string `query:"groupFilter"`
	Song        string `query:"songFilter"`
	ReleaseDate string `query:"releaseDateFilter"`
	Text        string `query:"textFilter"`
	Link        string `query:"linkFilter"`
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
}
