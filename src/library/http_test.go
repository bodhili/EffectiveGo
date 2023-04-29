package library

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHTTPRequestURI(t *testing.T) {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		t.Fatalf("Http server started is error")
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "host:[%q] \n", r.Host)
	_, _ = fmt.Fprintf(w, "method:[%q] \n", r.Method)
	_, _ = fmt.Fprintf(w, "proto:[%q] \n", r.Proto)
	_, _ = fmt.Fprintf(w, "remote-addr:[%q] \n", r.RemoteAddr)
	_, _ = fmt.Fprintf(w, "proto-major:[%q] \n", r.ProtoMajor)
	_, _ = fmt.Fprintf(w, "proto-minor:[%q] \n", r.ProtoMinor)
	_, _ = fmt.Fprintf(w, "request-uri:[%q] \n", r.RequestURI)
	_, _ = fmt.Fprintf(w, "*uri:[%q] \n", r.URL)

	_, _ = fmt.Fprintf(w, "header:[%q] \n", r.Header)
}
