package main

import "fmt"

// Coffee interface
type Coffee interface {
	Cost() int
}

// BasicCoffee is the concrete component
type BasicCoffee struct{}

func (c *BasicCoffee) Cost() int {
	return 5
}

// CoffeeDecorator is the decorator base class
type CoffeeDecorator struct {
	coffee Coffee
}

func (cd *CoffeeDecorator) Cost() int {
	return cd.coffee.Cost()
}

// Concrete decorators
type MilkDecorator struct {
	decorator CoffeeDecorator
}

func (md *MilkDecorator) Cost() int {
	return md.decorator.Cost() + 2
}

type SugarDecorator struct {
	decorator CoffeeDecorator
}

func (sd *SugarDecorator) Cost() int {
	return sd.decorator.Cost() + 1
}

type VanillaDecorator struct {
	decorator CoffeeDecorator
}

func (vd *VanillaDecorator) Cost() int {
	return vd.decorator.Cost() + 3
}

func main() {
	myCoffee := &BasicCoffee{}
	fmt.Printf("Cost of %T: $%d\n", myCoffee, myCoffee.Cost())

	coffeeWithMilk := &MilkDecorator{decorator: CoffeeDecorator{coffee: myCoffee}}
	fmt.Printf("Cost of %T: $%d\n", coffeeWithMilk, coffeeWithMilk.Cost())

	coffeeWithMilkAndSugar := &SugarDecorator{decorator: CoffeeDecorator{coffee: coffeeWithMilk}}
	fmt.Printf("Cost of %T: $%d\n", coffeeWithMilkAndSugar, coffeeWithMilkAndSugar.Cost())

	coffeeWithMilkSugarAndVanilla := &VanillaDecorator{decorator: CoffeeDecorator{coffee: coffeeWithMilkAndSugar}}
	fmt.Printf("Cost of %T: $%d\n", coffeeWithMilkSugarAndVanilla, coffeeWithMilkSugarAndVanilla.Cost())
}
