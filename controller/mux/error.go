package mux

type Error struct {
	Err        error
	StatusCode int
}

func (e Error) Error() string {
	return e.Err.Error()
}

func newError(code int, err error) error {
	return Error{Err: err, StatusCode: code}
}
