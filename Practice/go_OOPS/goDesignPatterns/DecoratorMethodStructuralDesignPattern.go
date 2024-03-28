package main

import "fmt"

type Beverage interface {
	Cost() float64
	Description() string
}

type Coffee struct{}

func (c *Coffee) Cost() float64 {
	return 5.00
}
func (c *Coffee) Description() string {
	return "Coffee"
}

//Decorator design pattern - Adding new functionality/more responsibility to existing component/struct/class dynamically without altering the existing component.
//Real use case -- Adding cheese topping to pizza, Adding extension to any component etc.

type MilkDecorator struct {
	beverage Beverage //Adding Beverage interface implemented "coffee" struct. Instead of this, If We can use struct composition with method overridding of cost() and description(), But there we should know the "coffee" struct cost and description value for maintaining value.
}

func (c *MilkDecorator) CostAndDescription() (float64, string) {
	return c.beverage.Cost() + 2, c.beverage.Description() + " Adding Milk" //Dynamically adding the values into Beverage interface implemented "coffee" struct. In future new structs can implement Beverage interface as well. Like Coffee struct we can have multiple structs as well which implements Beverage interface. Here MilkDecorator struct don't need to know the Coffee struct "cost" value as well.
}

func main() {
	coffee := &Coffee{}
	milkDecorator := MilkDecorator{beverage: coffee}
	fmt.Println("Coffee cost", coffee.Cost())
	fmt.Println("Coffee description --", coffee.Description())
	cost, description := milkDecorator.CostAndDescription()
	fmt.Println("Coffee with milk cost and description -- ", cost, "-----", description)
}
