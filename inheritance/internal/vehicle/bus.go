package vehicle

import "fmt"

type Bus struct {
	vehicleBase
	name string
}

func NewBus(name string) *Bus {
	b := &Bus{}
	b.name = name
	b.vehicleBase.wheelCount = 0
	return b
}

func (b *Bus) GetName() string {
	return b.name
}

func (b *Bus) SetName(name string) {
	b.name = name
}

func (b *Bus) GetType() Type {
	return BusType
}

func (b *Bus) ToString() string {
	str := fmt.Sprintf("Bus -> %s", b.GetName())
	return b.vehicleBase.ToString() + str
}
