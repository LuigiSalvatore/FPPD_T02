package main

//	O servidor só precisa gerar o mapa
//	Deve ser implementadas funções que modifiquem o mapa, movem os jogadores, e interajam com os elementos
//
//
//
import (
	"bufio"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

// Define os elementos do jogo
type Elemento struct {
	simbolo  rune
	cor      termbox.Attribute
	corFundo termbox.Attribute
	tangivel bool
}

// Personagem controlado pelo jogador
var personagem = Elemento{
	simbolo:  '☺',
	cor:      termbox.ColorBlack,
	corFundo: termbox.ColorDefault,
	tangivel: true,
}

// Parede
var parede = Elemento{
	simbolo:  '▤',
	cor:      termbox.ColorBlack | termbox.AttrBold | termbox.AttrDim,
	corFundo: termbox.ColorDarkGray,
	tangivel: true,
}

// Barrreira
var barreira = Elemento{
	simbolo:  '#',
	cor:      termbox.ColorRed,
	corFundo: termbox.ColorDefault,
	tangivel: true,
}

// Vegetação
var vegetacao = Elemento{
	simbolo:  '♣',
	cor:      termbox.ColorGreen,
	corFundo: termbox.ColorDefault,
	tangivel: false,
}

// Elemento vazio
var vazio = Elemento{
	simbolo:  ' ',
	cor:      termbox.ColorDefault,
	corFundo: termbox.ColorDefault,
	tangivel: false,
}

// Elemento para representar áreas não reveladas (efeito de neblina)
var neblina = Elemento{
	simbolo:  '.',
	cor:      termbox.ColorDefault,
	corFundo: termbox.ColorYellow,
	tangivel: false,
}

var mapa [][]Elemento
var posX, posY int
var ultimoElementoSobPersonagem = vazio
var lastElement_1 = vazio
var lastElement_2 = vazio
var lastElement_3 = vazio
var statusMsg string

var efeitoNeblina = false
var revelado [][]bool
var raioVisao int = 3

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	carregarMapa("mapa.txt")
}

func carregarMapa(nomeArquivo string) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		panic(err)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	y := 0
	for scanner.Scan() {
		linhaTexto := scanner.Text()
		var linhaElementos []Elemento
		var linhaRevelada []bool
		for x, char := range linhaTexto {
			elementoAtual := vazio
			switch char {
			case parede.simbolo:
				elementoAtual = parede
			case barreira.simbolo:
				elementoAtual = barreira
			case vegetacao.simbolo:
				elementoAtual = vegetacao
			case personagem.simbolo:
				// Atualiza a posição inicial do personagem
				posX, posY = x, y
				elementoAtual = vazio
			}
			linhaElementos = append(linhaElementos, elementoAtual)
			linhaRevelada = append(linhaRevelada, false)
		}
		mapa = append(mapa, linhaElementos)
		revelado = append(revelado, linhaRevelada)
		y++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func updatePos(novaPosX int, novaPosY int, elem Elemento) { // Cliente chama essa função para atualizar a posição do elemento

	if novaPosY >= 0 && novaPosY < len(mapa) && novaPosX >= 0 && novaPosX < len(mapa[novaPosY]) && mapa[novaPosY][novaPosX].tangivel == false {
		mapa[posY][posX] = lastElement_1         // Restaura o elemento anterior
		lastElement_1 = mapa[novaPosY][novaPosX] // Atualiza o elemento
		posX, posY = novaPosX, novaPosY          // Move o elemento
		mapa[posY][posX] = elem                  // Coloca o elemento na nova posição
	}
}
func mover(comando rune) {
	dx, dy := 0, 0
	switch comando {
	case 'w':
		dy = -1
	case 'a':
		dx = -1
	case 's':
		dy = 1
	case 'd':
		dx = 1
	}
	novaPosX, novaPosY := posX+dx, posY+dy
	if novaPosY >= 0 && novaPosY < len(mapa) && novaPosX >= 0 && novaPosX < len(mapa[novaPosY]) &&
		mapa[novaPosY][novaPosX].tangivel == false {
		mapa[posY][posX] = ultimoElementoSobPersonagem         // Restaura o elemento anterior
		ultimoElementoSobPersonagem = mapa[novaPosY][novaPosX] // Atualiza o elemento sob o personagem
		posX, posY = novaPosX, novaPosY                        // Move o personagem
		mapa[posY][posX] = personagem                          // Coloca o personagem na nova posição
	}
}

func interagir() {
	statusMsg = fmt.Sprintf("Interagindo em (%d, %d)", posX, posY)
}
