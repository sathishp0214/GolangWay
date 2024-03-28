package main

import (
	"fmt"
	"sync"
)

func main() {
	bag1 := Bag{10}
	bag2 := Bag{20}
	bag3 := Bag{30}

	customer := Customer{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		customer.SetOfBags = []Bag{bag1, bag2, bag3}
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Total weight", customer.CalculateWeight())
	}()

	wg.Wait()

	fmt.Println("completed")

}

type Customer struct {
	SetOfBags []Bag
}

type Bag struct {
	Weight int
}

func (c *Customer) CalculateWeight() int {
	sum := 0
	for _, i := range c.SetOfBags {
		sum = sum + i.Weight
	}
	return sum
}
