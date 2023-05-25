package vehicle

type Type int

const (
	BusType   Type = 0
	TrainType Type = 1
)

type Vehicle interface {
	GetType() Type

	GetName() string
	SetName(name string)

	GetWheelCount() int
	SetWheelCount(count int)

	ToString() string
}

type vehicleBase struct {
	wheelCount int
}

func (v *vehicleBase) SetWheelCount(count int) {
	v.wheelCount = count
}

func (v *vehicleBase) GetWheelCount() int {
	return v.wheelCount
}

func (v *vehicleBase) ToString() string {
	return "vehicle -> "
}
