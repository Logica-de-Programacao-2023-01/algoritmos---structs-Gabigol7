package main

import (
	"fmt"
	"time"
)

type Funcionário struct {
	nome    string
	salário float64
	idade   int
}

func (f *Funcionário) aumentarSalário(porcentagem float64) {
	f.salário *= (1 + (porcentagem / 100))
}

func (f *Funcionário) diminuirSalário(porcentagem float64) {
	f.salário *= (1 - (porcentagem / 100))
}

func (f *Funcionário) tempoDeServiço() int {
	idadeInício := 18
	anosDeServiço := time.Now().Year() - f.idade - idadeInício
	return anosDeServiço
}

func main() {
	funcionário := Funcionário{
		nome:    "João",
		salário: 3000.0,
		idade:   25,
	}

	fmt.Println("Salário inicial:", funcionário.salário)
	funcionário.aumentarSalário(10)
	fmt.Println("Salário após aumento:", funcionário.salário)
	funcionário.diminuirSalário(5)
	fmt.Println("Salário após diminuição:", funcionário.salário)

	fmt.Println("Tempo de serviço:", funcionário.tempoDeServiço(), "anos")
}
