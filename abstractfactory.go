package main

import "fmt"

// Abstract Factory Interface
type ComputerFactory interface {
	CreateCPU() CPU
	CreateRAM() RAM
}

// Concrete Factory for Laptop
type LaptopFactory struct{}

func (lf LaptopFactory) CreateCPU() CPU {
	return LaptopCPU{}
}

func (lf LaptopFactory) CreateRAM() RAM {
	return LaptopRAM{}
}

// Concrete Factory for Desktop
type DesktopFactory struct{}

func (df DesktopFactory) CreateCPU() CPU {
	return DesktopCPU{}
}

func (df DesktopFactory) CreateRAM() RAM {
	return DesktopRAM{}
}

// Abstract Product Interface for CPU
type CPU interface {
	Process() string
}

// Concrete Products for CPU
type LaptopCPU struct{}

func (lc LaptopCPU) Process() string {
	return "Laptop CPU is processing."
}

type DesktopCPU struct{}

func (dc DesktopCPU) Process() string {
	return "Desktop CPU is processing."
}

// Abstract Product Interface for RAM
type RAM interface {
	Storage() string
}

// Concrete Products for RAM
type LaptopRAM struct{}

func (lr LaptopRAM) Storage() string {
	return "Laptop RAM is storing data."
}

type DesktopRAM struct{}

func (dr DesktopRAM) Storage() string {
	return "Desktop RAM is storing data."
}


func AssembleComputer(factory ComputerFactory) string {
	cpu := factory.CreateCPU()
	ram := factory.CreateRAM()
	return fmt.Sprintf("Computer components: %s %s", cpu.Process(), ram.Storage())
}

func main() {
	laptopFactory := LaptopFactory{}
	desktopFactory := DesktopFactory{}

	laptop := AssembleComputer(laptopFactory)
	desktop := AssembleComputer(desktopFactory)

	fmt.Println("Assembling a laptop:")
	fmt.Println(laptop)

	fmt.Println("\nAssembling a desktop:")
	fmt.Println(desktop)
}