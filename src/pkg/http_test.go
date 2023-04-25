package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
)

type countHandler struct {
	mu sync.Mutex
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.n++

	_, err := fmt.Fprintf(w, "count is %d \n", h.n)
	if err != nil {
		fmt.Println("error")
		return
	}
}

func TestHttpClientHandle(t *testing.T) {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func TestHttpClientHandleFunc(t *testing.T) {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello from a HandleFunc #1!\n")
		if err != nil {
			return
		}
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello from a HandleFunc #2!\n")
		if err != nil {
			return
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		t.Log(err)
	}
}
