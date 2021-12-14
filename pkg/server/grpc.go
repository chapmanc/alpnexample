package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.

	"github.com/chapmanc/alpnexample/proto/go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunGRPC(ln net.Listener, tlsConfig *tls.Config) error {
	var opts []grpc.ServerOption
	tlsCredentials := credentials.NewTLS(tlsConfig)
	opts = append(opts, grpc.Creds(tlsCredentials))
	server := grpc.NewServer(opts...)
	helloworld.RegisterHelloWorldServiceServer(server, &helloWorldServiceServer{})
	go func() {
		if err := server.Serve(ln); err != nil {
			fmt.Printf("server returned: %v\n", err)
		}
	}()

	return nil
}

// petStoreServiceServer implements the PetStoreService API.
type helloWorldServiceServer struct {
	helloworld.UnimplementedHelloWorldServiceServer
}

// GetHello adds the pet associated with the given request into the PetStore.
func (s *helloWorldServiceServer) GetHello(ctx context.Context, req *helloworld.GetHelloRequest) (*helloworld.GetHelloResponse, error) {
	message := req.GetMsg()
	log.Println("Got a GetHelloWorldRequest", "message", message)

	return &helloworld.GetHelloResponse{Msg: fmt.Sprintf("Hello %v", message)}, nil
}
