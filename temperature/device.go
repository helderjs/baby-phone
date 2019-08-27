package temperature

type Device interface {
	ReadTemperature() float32
}
