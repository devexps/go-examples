package fruit

type Apple struct {
	name string
}

func NewApple(name string) *Apple {
	return &Apple{
		name: name,
	}
}

func (a *Apple) GetName() string {
	return a.name
}

func (a *Apple) SetName(name string) {
	a.name = name
}

func (a *Apple) GetType() Type {
	return AppleType
}
