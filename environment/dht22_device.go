package environment

import "github.com/d2r2/go-dht"

type dht22Device struct{}

func (d *dht22Device) ReadEnvironment() (float32, float32) {
	t, h, _, _ := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)

	return t, h
}

func (d *dht22Device) ReadTemperature() float32 {
	t, _ := d.ReadEnvironment()

	return t
}

func (d *dht22Device) ReadHumidity() float32 {
	_, h := d.ReadEnvironment()

	return h
}

func NewDht22Device() Device {
	return &dht22Device{}
}
