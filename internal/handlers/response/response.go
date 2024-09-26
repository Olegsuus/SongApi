package response

type AddSongResponse struct {
	ID int `json:"id"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}
