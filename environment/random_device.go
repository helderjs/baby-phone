package environment

import "math/rand"

type randomDevice struct{}

func (d *randomDevice) ReadEnvironment() (float32, float32) {
	return (rand.Float32() * 5) + 25, (rand.Float32() * 50) + 50
}

func (d *randomDevice) ReadTemperature() float32 {
	return (rand.Float32() * 5) + 25
}

func (d *randomDevice) ReadHumidity() float32 {
	return (rand.Float32() * 50) + 50
}

func NewRandomDevice() Device {
	return &randomDevice{}
}
