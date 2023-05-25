package fruit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApple(t *testing.T) {
	{
		var apple *Apple
		apple = NewApple("little")
		assert.NotNil(t, apple)
		assert.Equal(t, apple.GetName(), "little")
		assert.Equal(t, apple.GetType(), AppleType)
	}
	{
		var fruitApple Fruit
		fruitApple = NewFruit(AppleType, "middle")
		assert.NotNil(t, fruitApple)
		assert.Equal(t, fruitApple.GetName(), "middle")
		assert.Equal(t, fruitApple.GetType(), AppleType)

		fruitApple.SetName("big")
		assert.Equal(t, fruitApple.GetName(), "big")
	}
}

func TestBanana(t *testing.T) {
	{
		var banana *Banana
		banana = NewBanana("little")
		assert.NotNil(t, banana)
		assert.Equal(t, banana.GetName(), "little")
		assert.Equal(t, banana.GetType(), BananaType)

		banana.SetEnergy(int64(10))
		assert.Equal(t, banana.GetEnergy(), int64(10))
	}
	{
		var fruitBanana Fruit
		fruitBanana = NewFruit(BananaType, "middle")
		assert.NotNil(t, fruitBanana)
		assert.Equal(t, fruitBanana.GetName(), "middle")
		assert.Equal(t, fruitBanana.GetType(), BananaType)

		fruitBanana.SetName("big")
		assert.Equal(t, fruitBanana.GetName(), "big")
	}
}

func TestNewFruit(t *testing.T) {
	{
		unknown := NewFruit(3, "unknown")
		assert.Nil(t, unknown)
	}
	{
		apple := NewFruit(AppleType, "big")
		assert.NotNil(t, apple)
		assert.Equal(t, apple.GetName(), "big")
		assert.Equal(t, apple.GetType(), AppleType)

		apple.SetName("small")
		assert.Equal(t, apple.GetName(), "small")
	}
	{
		banana := NewFruit(BananaType, "big")
		assert.NotNil(t, banana)
		assert.Equal(t, banana.GetName(), "big")
		assert.Equal(t, banana.GetType(), BananaType)

		banana.SetName("small")
		assert.Equal(t, banana.GetName(), "small")
	}
}
