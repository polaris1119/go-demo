package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", logger(index))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type Response struct {
	http.ResponseWriter
	Status    int
	Size      int64
	Committed bool
}

func (r *Response) WriteHeader(code int) {
	if r.Committed {
		return
	}
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
	r.Committed = true
}

func (r *Response) Write(b []byte) (n int, err error) {
	if !r.Committed {
		if r.Status == 0 {
			r.Status = http.StatusOK
		}
		r.WriteHeader(r.Status)
	}
	n, err = r.ResponseWriter.Write(b)
	r.Size += int64(n)
	return
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := &Response{
			ResponseWriter: w,
		}

		start := time.Now()
		defer func() {
			elapse := time.Now().Sub(start)

			log.Printf("addr:%s uri:%s method:%s status:%d size:%d delay:%d\n",
				req.RemoteAddr,
				req.RequestURI,
				req.Method,
				resp.Status,
				resp.Size,
				elapse,
			)
		}()

		next(resp, req)
	}
}
