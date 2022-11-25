package srv

import "net/http"

type StatusWriter struct {
	w           http.ResponseWriter
	statusCode  int
	codeWritten bool
}

func NewStatusWriter(w http.ResponseWriter) *StatusWriter {
	return &StatusWriter{
		w: w,
	}
}

func (sw *StatusWriter) Header() http.Header {
	return sw.w.Header()
}

func (sw *StatusWriter) Write(p []byte) (n int, err error) {
	return sw.w.Write(p)
}

func (sw *StatusWriter) Code() int {
	return sw.statusCode
}

func (sw *StatusWriter) WriteHeader(code int) {
	if sw.codeWritten {
		return
	}
	sw.statusCode = code
	sw.w.WriteHeader(code)
	sw.codeWritten = true
}
