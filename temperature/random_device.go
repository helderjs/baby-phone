package temperature

import "math/rand"

type randomDevice struct{}

func (d *randomDevice) ReadTemperature() float32 {
	return (rand.Float32() * 5) + 25
}

func NewRandomDevice() Device {
	return &randomDevice{}
}
