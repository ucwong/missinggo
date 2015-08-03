package missinggo

import (
	"io"
	"net/http"
)

type StatusResponseWriter struct {
	RW           http.ResponseWriter
	Code         int
	BytesWritten int64
}

var _ http.ResponseWriter = &StatusResponseWriter{}

func (me *StatusResponseWriter) Header() http.Header {
	return me.RW.Header()
}

func (me *StatusResponseWriter) Write(b []byte) (n int, err error) {
	if me.Code == 0 {
		me.Code = 200
	}
	n, err = me.RW.Write(b)
	me.BytesWritten += int64(n)
	return
}

func (me *StatusResponseWriter) WriteHeader(code int) {
	me.RW.WriteHeader(code)
	me.Code = code
}

type ReaderFromStatusResponseWriter struct {
	StatusResponseWriter
	io.ReaderFrom
}

func NewReaderFromStatusResponseWriter(w http.ResponseWriter) *ReaderFromStatusResponseWriter {
	return &ReaderFromStatusResponseWriter{
		StatusResponseWriter{RW: w},
		w.(io.ReaderFrom),
	}
}