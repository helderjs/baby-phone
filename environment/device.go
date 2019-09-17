package environment

type Device interface {
	ReadTemperature() float32
	ReadHumidity() float32
	ReadEnvironment() (float32, float32)
}
