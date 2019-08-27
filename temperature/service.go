package temperature

import "errors"

var ErrInvalidArgument = errors.New("invalid argument")

type Service interface {
	GetTemperature() float32
}

type service struct {
	device Device
}

func (s *service) GetTemperature() float32 {
	return s.device.ReadTemperature()
}

func NewService(d Device) Service {
	return &service{device: d}
}
