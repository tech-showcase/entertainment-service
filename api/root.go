package api

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func Activate(port int) {
	gRPCServer := grpc.NewServer()

	RegisterMovieGRPCAPI(gRPCServer)

	lisAddress := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", lisAddress)
	if err != nil {
		panic(err)
	}

	err = gRPCServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
