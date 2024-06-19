package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"

	"github.com/nsf/termbox-go"
)

var mutex sync.Mutex

type Elemento struct {
	Simbolo  rune
	Cor      termbox.Attribute
	CorFundo termbox.Attribute
	Tangivel bool
}
type Jogador struct {
	ID      int
	Element Elemento
	posX    int
	posY    int
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
	player := new(Jogador)
	desenhaTudo()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return // Sair do programa
			}
			if ev.Ch == 'w' || ev.Ch == 'a' || ev.Ch == 's' || ev.Ch == 'd' || ev.Ch == 'e' {
				err = client.Call("Servidor.ListenInput", ev.Ch, &player)
				if err != nil {
					fmt.Println("Erro ao enviar comando:", err)
				}
			}
			err = client.Call("Servidor.SendMapa", maquina, &mapa)
			if err != nil {
				fmt.Println("Erro ao obter mapa:", err)
			}
			desenhaTudo()
		}
	}
}
