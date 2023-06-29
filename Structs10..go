package main

import (
	"fmt"
)

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(index int) {
	if index >= 0 && index < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:index], f.avaliacoes[index+1:]...)
	}
}

func (f *Filme) calcularMediaAvaliacoes() float64 {
	soma := 0
	for _, nota := range f.avaliacoes {
		soma += nota
	}
	media := float64(soma) / float64(len(f.avaliacoes))
	return media
}

func (f *Filme) imprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.titulo)
	fmt.Printf("Diretor: %s\n", f.diretor)
	fmt.Printf("Ano: %d\n", f.ano)
	fmt.Printf("Avaliações: %v\n", f.avaliacoes)
	fmt.Printf("Média das Avaliações: %.2f\n", f.calcularMediaAvaliacoes())
}

func main() {
	// Exemplo de uso da struct Filme
	filme := Filme{
		titulo:  "Matrix",
		diretor: "Lana Wachowski",
		ano:     1999,
	}

	filme.adicionarAvaliacao(8)
	filme.adicionarAvaliacao(9)
	filme.adicionarAvaliacao(7)
	filme.adicionarAvaliacao(10)

	filme.imprimirInformacoes()

	filme.removerAvaliacao(1)

	filme.imprimirInformacoes()
}
