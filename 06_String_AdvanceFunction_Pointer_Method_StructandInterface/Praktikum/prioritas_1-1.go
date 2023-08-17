package main

import "fmt"

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) CalculateDistance() float64 {
	return c.fuelin * 1.5
}

func main() {
	car := Car{"SUV", 10.5}
	perkiraanPengeluaran := car.CalculateDistance()
	fmt.Printf("car type: %s , est. distance: %.2f\n", car.carType, perkiraanPengeluaran)
}
