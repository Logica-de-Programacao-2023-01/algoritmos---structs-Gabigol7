package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) removerNota(index int) {
	if index >= 0 && index < len(a.notas) {
		a.notas = append(a.notas[:index], a.notas[index+1:]...)
	}
}

func (a *Aluno) calcularMedia() float64 {
	soma := 0.0
	for _, nota := range a.notas {
		soma += nota
	}
	return soma / float64(len(a.notas))
}

func (a *Aluno) imprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Média das notas: %.2f\n", a.calcularMedia())
}

func main() {
	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{8.5, 7.2, 9.0},
	}

	aluno.imprimirInformacoes()

	aluno.adicionarNota(6.8)
	aluno.imprimirInformacoes()

	aluno.removerNota(1)
	aluno.imprimirInformacoes()
}
