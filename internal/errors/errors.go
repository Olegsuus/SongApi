package errors

type AppError interface {
	error
	ReqErrorStatus() int
	UErrorText() string
}

type ReqError struct {
	Status int    `json:"status"`
	Text   string `json:"text"`
}

func (e ReqError) Error() string {
	return e.Text
}

func (e ReqError) ReqErrorStatus() int {
	return e.Status
}

func (e ReqError) UErrorText() string {
	return e.Text
}

func NewReqError(status int, text string) ReqError {
	return ReqError{
		Status: status,
		Text:   text,
	}
}
