package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"time"
)

func ApplyTracerClient(operationName string, endpoint endpoint.Endpoint, tracer stdopentracing.Tracer) (wrappedEndpoint endpoint.Endpoint) {
	wrappedEndpoint = opentracing.TraceClient(tracer, operationName)(endpoint)
	return
}

func ApplyLogger(operationName string, nextEndpoint endpoint.Endpoint, logger log.Logger) (wrappedEndpoint endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		defer func(begin time.Time) {
			logger.Log("operationName", operationName, "elapsedTime", time.Since(begin))
		}(time.Now())

		return nextEndpoint(ctx, request)
	}
}
