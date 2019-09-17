package environment

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/examples/shipping/cargo"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(s Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getEnvironmentHandler := kithttp.NewServer(
		makeGetEnvironmentEndpoint(s),
		decodeGetEnvironmentRequest,
		encodeResponse,
		opts...,
	)
	getTemperatureHandler := kithttp.NewServer(
		makeGetTemperatureEndpoint(s),
		decodeGetTemperatureRequest,
		encodeResponse,
		opts...,
	)
	getHumidityHandler := kithttp.NewServer(
		makeGetHumidityEndpoint(s),
		decodeGetHumidityRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/v1/environment", getEnvironmentHandler).Methods("GET")
	r.Handle("/v1/environment/temperature", getTemperatureHandler).Methods("GET")
	r.Handle("/v1/environment/humidity", getHumidityHandler).Methods("GET")

	return r
}

func decodeGetEnvironmentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getEnvironmentRequest{}, nil
}

func decodeGetTemperatureRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getTemperatureRequest{}, nil
}

func decodeGetHumidityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getHumidityRequest{}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
