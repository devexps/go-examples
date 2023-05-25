package human

type Mother struct {
}

func NewMother() *Mother {
	return &Mother{}
}

func (m *Mother) GetName() string {
	return "HienNTT"
}

func (m *Mother) Say() string {
	return "I am " + m.GetName()
}
