package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func ImprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço completo de %s:\n", p.Nome)
	fmt.Printf("Rua: %s, Número: %d\n", p.Endereco.Rua, p.Endereco.Numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", p.Endereco.Cidade, p.Endereco.Estado)
}

func main() {
	p := Pessoa{
		Nome:  "João",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua A",
			Numero: 123,
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}

	ImprimirEnderecoCompleto(p)
}
