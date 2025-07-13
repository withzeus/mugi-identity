package lib

type HttpError interface {
	error
	Status() int
}

type HttpStatusError struct {
	Code int
	Err  error
}

func (se HttpStatusError) Error() string {
	return se.Err.Error()
}

func (se HttpStatusError) Status() int {
	return se.Code
}
