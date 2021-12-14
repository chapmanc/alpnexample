package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/chapmanc/alpnexample/pkg/secure"
	"github.com/chapmanc/alpnexample/pkg/server"
	"golang.org/x/net/http2"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:7999")
	if err != nil {
		fmt.Printf("error starting listener: %v\n", err)
		os.Exit(1)
	}
	certs := secure.GetCerts()
	certPool := secure.NewCertPool(certs)
	h2TLSConfig := secure.GetTlsConfig(certs, certPool, []string{"h2", "grpc", "foo"})
	tlsLn := tls.NewListener(ln, h2TLSConfig)

	mux := http.NewServeMux()
	mux.HandleFunc("/h2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("handling /h2 for client\n")
		w.Write([]byte("gotcha"))
	})

	shutdownCh := make(chan struct{})
	sighupCh := make(chan os.Signal, 4)
	signal.Notify(sighupCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sighupCh
		close(shutdownCh)
	}()

	server.RunGRPC(ln, h2TLSConfig)

	server := &http.Server{
		Addr:    "127.0.0.1:7999",
		Handler: mux,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){
			"foo": server.GetFooProtocol(shutdownCh),
		},
	}

	http2.ConfigureServer(server, nil)

	go func() {
		err = server.Serve(tlsLn)
		fmt.Printf("server returned: %v\n", err)
	}()

	select {
	case <-shutdownCh:
		return
	}
}
