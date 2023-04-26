package lan

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestEchoUrl(t *testing.T) {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch:%v \n", err)
			os.Exit(1)
		}
		bs, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}

		fmt.Printf("body:[%s]", bs)
	}
}

func TestMulti(t *testing.T) {
	start := time.Now()
	ch := make(chan string, 100)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	bs, err := io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, bs, url)
}
