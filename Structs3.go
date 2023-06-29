package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func CalcularArea(triangulo Triangulo) float64 {
	area := triangulo.base * triangulo.altura / 2
	return area
}

func main() {
	triangulo := Triangulo{base: 10, altura: 5}
	area := CalcularArea(triangulo)
	fmt.Println("A área do triângulo é:", area)
}
