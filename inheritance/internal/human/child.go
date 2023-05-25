package human

type Child struct {
	Father
	Mother
}

func NewChild() *Child {
	return &Child{}
}

func (c *Child) GetName() string {
	return "BaoNG"
}

func (c *Child) Say() string {
	return "I am " + c.GetName() + ", My Father is " + c.Father.GetName() + ", My Mother is " + c.Mother.GetName()
}
