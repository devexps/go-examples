package human

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChild(t *testing.T) {
	father := NewFather()
	assert.NotNil(t, father)

	mother := NewMother()
	assert.NotNil(t, mother)

	child := NewChild()
	assert.NotNil(t, child)

	fmt.Println(father.Say())
	fmt.Println(mother.Say())
	fmt.Println(child.Say())
}
