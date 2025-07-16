package lib

import "net/http"

type RestServeMux struct {
	*http.ServeMux
}

func NewRestServeMux() *RestServeMux {
	return &RestServeMux{http.NewServeMux()}
}

func (rx *RestServeMux) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	rx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

func (rx *RestServeMux) Post(path string, handler func(http.ResponseWriter, *http.Request) error) {
	rx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := handler(w, r)
			if err, throwse := err.(StatusError); throwse {
				rx.WriteStatusError(w, err)
			}
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

func (rx *RestServeMux) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	rx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

func (rx *RestServeMux) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	rx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

func (rx *RestServeMux) WriteStatusError(w http.ResponseWriter, err StatusError) {
	http.Error(w, http.StatusText(err.Status()), http.StatusMethodNotAllowed)
}

type StatusError interface {
	error
	Status() int
}

type HttpStatusCode struct {
	Code int
	Msg  string
}

func NewHttpStatusCode(code int, msg string) *HttpStatusCode {
	return &HttpStatusCode{Code: code, Msg: msg}
}

func (se HttpStatusCode) Error() string {
	return se.Msg
}

func (se HttpStatusCode) Status() int {
	return se.Code
}
