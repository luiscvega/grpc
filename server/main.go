package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/luiscvega/grpc/x"

	"google.golang.org/grpc"
)

type Resizer struct{}

func (r Resizer) ResizeImage(ctx context.Context, image *x.Image) (*x.Error, error) {
	fmt.Println(len(image.Bytes))

	return &x.Error{Text: "It is done!"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	x.RegisterResizerServer(s, Resizer{})

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
