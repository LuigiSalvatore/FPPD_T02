package main

import (
	"fmt"
	"net/rpc"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Elemento struct {
	Simbolo  rune
	Cor      termbox.Attribute
	CorFundo termbox.Attribute
	Tangivel bool
}
type Jogador struct {
	ID      int
	Element Elemento
	PosX    int
	PosY    int
	Online  bool
}

var personagem_1 = Elemento{
	Simbolo:  '☺',
	Cor:      termbox.ColorRed,
	CorFundo: termbox.ColorDefault,
	Tangivel: true,
}
var personagem_2 = Elemento{
	Simbolo:  '☺',
	Cor:      termbox.ColorGreen,
	CorFundo: termbox.ColorDefault,
	Tangivel: true,
}
var personagem_3 = Elemento{
	Simbolo:  '☺',
	Cor:      termbox.ColorBlue,
	CorFundo: termbox.ColorDefault,
	Tangivel: true,
}

// Parede
var parede = Elemento{
	Simbolo:  '▤',
	Cor:      termbox.ColorYellow | termbox.AttrBold | termbox.AttrDim,
	CorFundo: termbox.ColorDarkGray,
	Tangivel: true,
}

// Barrreira
var barreira = Elemento{
	Simbolo:  '#',
	Cor:      termbox.ColorRed,
	CorFundo: termbox.ColorDefault,
	Tangivel: true,
}

// Vegetação
var vegetacao = Elemento{
	Simbolo:  '♣',
	Cor:      termbox.ColorGreen,
	CorFundo: termbox.ColorDefault,
	Tangivel: false,
}

// Elemento vazio
var vazio = Elemento{
	Simbolo:  ' ',
	Cor:      termbox.ColorDefault,
	CorFundo: termbox.ColorDefault,
	Tangivel: false,
}

// Elemento para representar áreas não reveladas (efeito de neblina)
var neblina = Elemento{
	Simbolo:  '.',
	Cor:      termbox.ColorDefault,
	CorFundo: termbox.ColorYellow,
	Tangivel: false,
}

var mapa [][]Elemento
var statusMsg string
var player Jogador

func desenhaTudo() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y, linha := range mapa {
		for x, elem := range linha {
			termbox.SetCell(x, y, elem.Simbolo, elem.Cor, elem.CorFundo)
		}
	}

	desenhaBarraDeStatus()

	termbox.Flush()
}
func desenhaBarraDeStatus() {
	for i, c := range statusMsg {
		termbox.SetCell(i, len(mapa)+1, c, termbox.ColorBlack, termbox.ColorDefault)
	}
	msg := "Use WASD para mover e E para interagir. ESC para sair."
	for i, c := range msg {
		termbox.SetCell(i, len(mapa)+3, c, termbox.ColorBlack, termbox.ColorDefault)
	}
}
func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	if len(os.Args) != 2 {
		fmt.Println("Uso:", os.Args[0], " <maquina>")
	} else {
		fmt.Println("Conectando a", os.Args[1])
	}

	porta := 8973
	maquina := os.Args[1]
	var id int

	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	err = client.Call("Servidor.SendMapa", maquina, &mapa)
	if err != nil || len(mapa) == 0 {
		fmt.Println("Erro ao obter mapa:", err)
		return
	}
	err = client.Call("Servidor.GetID", maquina, &id)
	if err != nil {
		fmt.Println("Erro ao obter ID:", err)
		return
	}
	fmt.Println("ID:", id)
	// err = client.Call("Servidor.GetPlayer", maquina, &player)
	// if err != nil {
	// 	fmt.Println("Erro ao obter jogador:", err)
	// 	return
	// } else {
	// 	err = client.Call("Servidor.AckPlayer", maquina, &player)
	// 	if err != nil {
	// 		fmt.Println("Erro ao enviar jogador:", err)
	// 		return
	// 	}
	// }
	go pega_e_desenha_tudo_do_mapa_porra(client, maquina)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return // Sair do programa
			}
			if ev.Ch == 'w' || ev.Ch == 'a' || ev.Ch == 's' || ev.Ch == 'd' || ev.Ch == 'e' {
				err = client.Call("Servidor.ListenInput", ev.Ch, &id)
				if err != nil {
					fmt.Println("Erro ao enviar comando:", err)
				}
			}
			err = client.Call("Servidor.SendMapa", maquina, &mapa)
			if err != nil {
				fmt.Println("Erro ao obter mapa:", err)
			}
		}
	}
}
func pega_e_desenha_tudo_do_mapa_porra(client *rpc.Client, maquina string) {
	for {
		err := client.Call("Servidor.SendMapa", maquina, &mapa)
		if err != nil {
			fmt.Println("Erro ao obter mapa:", err)
		}
		desenhaTudo()
		time.Sleep(100 * time.Millisecond)

	}
}
