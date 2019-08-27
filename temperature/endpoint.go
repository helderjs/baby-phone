package temperature

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type getTemperatureRequest struct{}

type getTemperatureResponse struct {
	Temperature string `json:"temperature"`
}

func makeGetTemperatureEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(getTemperatureRequest)
		t := s.GetTemperature()

		return getTemperatureResponse{Temperature: fmt.Sprintf("%2.2f", t)}, nil
	}
}
