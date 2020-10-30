package api

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RegisterObservabilityHTTPAPI(mux *http.ServeMux) {
	mux.Handle("/observability/metrics", promhttp.Handler())
}
