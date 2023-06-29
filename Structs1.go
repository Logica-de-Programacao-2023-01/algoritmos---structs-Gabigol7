package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calcularArea(c Circulo) float64 {
	area := math.Pi * c.raio * c.raio
	return area
}

func main() {
	circulo := Circulo{raio: 5.0}
	area := calcularArea(circulo)
	fmt.Println("A área do círculo é:", area)
}
