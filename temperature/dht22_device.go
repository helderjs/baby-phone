package temperature

import "github.com/d2r2/go-dht"

type dht22Device struct{}

func (d *dht22Device) ReadTemperature() float32 {
	temperature, _, _, _ := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)

	return temperature
}

func NewDht22Device() Device {
	return &dht22Device{}
}
