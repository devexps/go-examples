package fruit

type Banana struct {
	name   string
	energy int64
}

func NewBanana(name string) *Banana {
	return &Banana{
		name:   name,
		energy: 0,
	}
}

func (b *Banana) GetName() string {
	return b.name
}

func (b *Banana) SetName(name string) {
	b.name = name
}

func (b *Banana) GetType() Type {
	return BananaType
}

func (b *Banana) GetEnergy() int64 {
	return b.energy
}

func (b *Banana) SetEnergy(energy int64) {
	b.energy = energy
}
