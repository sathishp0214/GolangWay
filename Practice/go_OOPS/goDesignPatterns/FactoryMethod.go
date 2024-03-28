package main

import "fmt"

type Animal interface {
	MakeNoise() string
}

type Dog struct{}

func (d *Dog) MakeNoise() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) MakeNoise() string {
	return "Meow!"
}

//Factoy method name got - Because as users We got the products and not awares how its made inside a factory.

//Factory Method Design Pattern - Hides the objects creation logic from users/clients-- Instead of directly creating objects for structs, Creating objects hiddenly with a intermediate seperate below function by returning interface type for different structs with respect to different inputs.

//USE CASE - For creating objects for users/clients for any purpose through interface of different structs/classes
func FactoryMethod(t string) Animal {
	switch t {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}

func main() {

	AnimalObj := FactoryMethod("dog")
	fmt.Println(AnimalObj.MakeNoise())

	AnimalObj = FactoryMethod("cat")
	fmt.Println(AnimalObj.MakeNoise())
}
