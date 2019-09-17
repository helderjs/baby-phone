package environment

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type getEnvironmentRequest struct{}

type getEnvironmentResponse struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
}

type getTemperatureRequest struct{}

type getTemperatureResponse struct {
	Temperature string `json:"temperature"`
}

type getHumidityRequest struct{}

type getHumidityResponse struct {
	Humidity string `json:"humidity"`
}

func makeGetEnvironmentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(getEnvironmentRequest)
		t, h := s.GetEnvironment()

		return getEnvironmentResponse{
			Temperature: fmt.Sprintf("%2.2f", t),
			Humidity:    fmt.Sprintf("%2.2f", h),
		}, nil
	}
}

func makeGetTemperatureEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(getTemperatureRequest)
		t := s.GetTemperature()

		return getTemperatureResponse{Temperature: fmt.Sprintf("%2.2f", t)}, nil
	}
}

func makeGetHumidityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(getHumidityRequest)
		h := s.GetHumidity()

		return getHumidityResponse{Humidity: fmt.Sprintf("%2.2f", h)}, nil
	}
}
