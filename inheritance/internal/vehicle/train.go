package vehicle

import "fmt"

type Train struct {
	vehicleBase
	name string
}

func NewTrain(name string) *Train {
	t := &Train{}
	t.name = name
	t.vehicleBase.wheelCount = 0
	return t
}

func (t *Train) GetName() string {
	return t.name
}

func (t *Train) SetName(name string) {
	t.name = name
}

func (t *Train) GetType() Type {
	return TrainType
}

func (t *Train) ToString() string {
	str := fmt.Sprintf("Train -> %s - %d", t.GetName(), t.GetWheelCount())
	return t.vehicleBase.ToString() + str
}
