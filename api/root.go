package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

func Activate(port int) {
	lisAddress := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", lisAddress)
	if err != nil {
		panic(err)
	}

	cMux := cmux.New(listener)
	ActivateGRPC(cMux)
	ActivateHTTP(cMux)

	cMux.Serve()
}

func ActivateGRPC(cMux cmux.CMux) {
	gRPCServer := grpc.NewServer()
	RegisterMovieGRPCAPI(gRPCServer)

	grpcListener := cMux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	go func() {
		err := gRPCServer.Serve(grpcListener)
		if err != nil {
			panic(err)
		}
	}()
}

func ActivateHTTP(cMux cmux.CMux) {
	mux := http.NewServeMux()
	RegisterObservabilityHTTPAPI(mux)

	httpServer := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, mux),
	}

	httpListener := cMux.Match(cmux.HTTP1())
	go func() {
		err := httpServer.Serve(httpListener)
		if err != nil {
			panic(err)
		}
	}()
}
