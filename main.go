package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Wheel int
}

type Truck struct {
	BedSize int
	Car     Car
}

func main() {
	//nested struct
	laneTruck := Truck{
		BedSize: 10,
		Car: Car{
			Brand: "Toyota",
			Model: "Camry",
			Wheel: 10,
		},
	}

	fmt.Println("Car Brand:", laneTruck.Car.Brand)
	fmt.Println("Car Model:", laneTruck.Car.Model)
	fmt.Println("Total Wheel:", laneTruck.Car.Wheel)
	fmt.Println("Bedsize:", laneTruck.BedSize, "cm")

	//struct as a method
	r := rectangle{12, 20}
	fmt.Println(r.area())

}

// struct as a method
type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int {
	return r.width * r.height
}
