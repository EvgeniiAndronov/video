package videoserver

import "net/http"

type responceWriter struct {
	http.ResponseWriter
	code int
}

func (w *responceWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}
