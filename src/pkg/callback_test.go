package pkg

import (
	"context"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestCallBack(t *testing.T) {
	err := listenAndServe(":8080", http.DefaultServeMux, func(addr string, maxHeaderBytes int) {
		t.Logf("listener successfully, and addr:[%v] maxHeaderBytes:[%d]\n", addr, maxHeaderBytes)
	})

	if err != nil {
		t.Fatal(err)
	}

	hook(t, nil, err)
}

type listener func(addr string, maxHeaderBytes int) // defined *

func listenAndServe(addr string, handler http.Handler, listener listener) error {
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	listen, err := net.Listen("tcp", server.Addr)

	if err != nil {
		return err
	}

	listener(server.Addr, server.MaxHeaderBytes)
	return server.Serve(listen)
}

func hook(t *testing.T, server *http.Server, startErr error) {
	sigChan := make(chan os.Signal, 1)
	go func() {
		sig := <-sigChan
		t.Logf("received signal %v, stopping server...\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			t.Logf("listener failure, and addr:[%v] maxHeaderBytes:[%d]\n", server.Addr, server.MaxHeaderBytes)
		}

		if startErr != nil && startErr != http.ErrServerClosed {
			t.Logf("failed to start http server,and addr:[%v] maxHeaderBytes:[%d] \n", server.Addr, server.MaxHeaderBytes)
		}
	}()
}
