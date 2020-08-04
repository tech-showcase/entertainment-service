package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type (
	GRPCEndpoint struct {
		Endpoint endpoint.Endpoint
		Decoder  grpctransport.DecodeRequestFunc
		Encoder  grpctransport.EncodeResponseFunc
	}
)
