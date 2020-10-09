package middleware

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
)

func ApplyTracerClient(operationName string, endpoint endpoint.Endpoint, tracer stdopentracing.Tracer) (wrappedEndpoint endpoint.Endpoint) {
	wrappedEndpoint = opentracing.TraceClient(tracer, operationName)(endpoint)
	return
}
