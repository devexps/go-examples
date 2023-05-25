package main

import (
	"fmt"
	"github.com/devexps/go-examples/inheritance/internal/fruit"
	"github.com/devexps/go-examples/inheritance/internal/human"
	"github.com/devexps/go-examples/inheritance/internal/vehicle"
)

func main() {
	var apple *fruit.Apple
	apple = fruit.NewApple("little")
	fmt.Println(apple.GetType(), apple.GetName())

	var fruitApple fruit.Fruit
	fruitApple = fruit.NewFruit(fruit.AppleType, "middle")
	fmt.Println(fruitApple.GetType(), fruitApple.GetName())

	fruitApple.SetName("big")
	fmt.Println(fruitApple.GetType(), fruitApple.GetName())

	var banana *fruit.Banana
	banana = fruit.NewBanana("little")
	banana.SetEnergy(10)
	fmt.Println(banana.GetType(), banana.GetName(), banana.GetEnergy())

	var fruitBanana fruit.Fruit
	fruitBanana = fruit.NewFruit(fruit.BananaType, "middle")
	fmt.Println(fruitBanana.GetType(), fruitBanana.GetName())

	fruitBanana.SetName("big")
	fmt.Println(fruitBanana.GetType(), fruitBanana.GetName())

	var train *vehicle.Train
	train = vehicle.NewTrain("little")
	train.SetWheelCount(100)
	fmt.Println(train.GetType(), train.GetName(), train.GetWheelCount(), train.ToString())

	var trainVehicle vehicle.Vehicle
	trainVehicle = vehicle.Vehicle(train)
	fmt.Println(trainVehicle.GetType(), trainVehicle.GetName(), trainVehicle.GetWheelCount(), trainVehicle.ToString())

	trainVehicle.SetWheelCount(200)
	fmt.Println(trainVehicle.GetType(), trainVehicle.GetName(), trainVehicle.GetWheelCount(), trainVehicle.ToString())

	var bus *vehicle.Bus
	bus = vehicle.NewBus("little")
	fmt.Println(bus.GetType(), bus.GetName(), bus.GetWheelCount(), bus.ToString())

	var busVehicle vehicle.Vehicle
	busVehicle = vehicle.NewBus("big")
	fmt.Println(busVehicle.GetType(), busVehicle.GetName(), busVehicle.GetWheelCount(), busVehicle.ToString())

	busVehicle.SetName("middle")
	busVehicle.SetWheelCount(100)
	fmt.Println(busVehicle.GetType(), busVehicle.GetName(), busVehicle.GetWheelCount(), busVehicle.ToString())

	father := human.NewFather()
	mother := human.NewMother()
	child := human.NewChild()
	fmt.Println(father.Say())
	fmt.Println(mother.Say())
	fmt.Println(child.Say())
}
