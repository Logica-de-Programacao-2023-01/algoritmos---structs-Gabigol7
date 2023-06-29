package main

import "fmt"

type Playlist struct {
	nome    string
	titulos []string
}

func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	resultado := []Playlist{}

	for _, playlist := range playlists {
		for _, musica := range playlist.titulos {
			if musica == titulo {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	playlists := []Playlist{
		Playlist{
			nome: "Favoritos",
			titulos: []string{
				"Musica1",
				"Musica2",
				"Musica3",
			},
		},
		Playlist{
			nome: "Rock",
			titulos: []string{
				"Musica4",
				"Musica5",
				"Musica6",
			},
		},
		Playlist{
			nome: "Pop",
			titulos: []string{
				"Musica1",
				"Musica7",
				"Musica8",
			},
		},
	}

	tituloBuscado := "Musica1"
	resultado := buscarPlaylistsPorTitulo(tituloBuscado, playlists)

	if len(resultado) == 0 {
		fmt.Println("Nenhuma playlist encontrada com o título:", tituloBuscado)
	} else {
		fmt.Println("Playlists encontradas com o título", tituloBuscado+":")
		for _, playlist := range resultado {
			fmt.Println(playlist.nome)
		}
	}
}
