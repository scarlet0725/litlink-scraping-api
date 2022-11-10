package model

type AppError struct {
	Msg  string
	Code int
}

func (e *AppError) Error() string {
	return e.Msg
}
