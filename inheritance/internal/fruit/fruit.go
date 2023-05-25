package fruit

type Type int

const (
	AppleType  Type = 0
	BananaType Type = 1
)

type Fruit interface {
	GetName() string
	SetName(name string)

	GetType() Type
}

func NewFruit(t Type, name string) Fruit {
	switch t {
	case AppleType:
		return NewApple(name)
	case BananaType:
		return NewBanana(name)
	}
	return nil
}
