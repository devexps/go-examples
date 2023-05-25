package human

type Father struct {
}

func NewFather() *Father {
	return &Father{}
}

func (f *Father) GetName() string {
	return "ThangN"
}

func (f *Father) Say() string {
	return "I am " + f.GetName()
}
