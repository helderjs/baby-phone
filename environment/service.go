package environment

import "errors"

var ErrInvalidArgument = errors.New("invalid argument")

type Service interface {
	GetTemperature() float32
	GetHumidity() float32
	GetEnvironment() (float32, float32)
}

type service struct {
	device Device
}

func (s *service) GetEnvironment() (float32, float32) {
	return s.device.ReadEnvironment()
}

func (s *service) GetTemperature() float32 {
	return s.device.ReadTemperature()
}

func (s *service) GetHumidity() float32 {
	return s.device.ReadHumidity()
}

func NewService(d Device) Service {
	return &service{device: d}
}
