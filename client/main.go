package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/luiscvega/grpc/x"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := x.NewResizerClient(conn)

	bs, err := ioutil.ReadFile("/home/luis/Desktop/photo.jpeg")
	if err != nil {
		panic(err)
	}

	e, err := c.ResizeImage(context.TODO(), &x.Image{Bytes: bs})
	if err != nil {
		panic(err)
	}

	log.Printf("Error: %s", e.Text)
}
