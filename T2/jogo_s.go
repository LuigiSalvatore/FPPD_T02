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

func connect() {
	// Conectar ao servidor
	//cria um objeto jogador( ID , ELEMENT )
}

// Define os elementos do jogo
type Elemento struct {
	simbolo  rune
	cor      termbox.Attribute
	corFundo termbox.Attribute
	tangivel bool
}
type Jogador struct {
	ID      int
	Element Elemento
	TX      int
	RX      int
	jogador.posX    int
	jogador.posY    int
	Online  bool
}
type Servidor struct {
	Jogadores [3]Jogador
}

var personagem_1 = Elemento{
	simbolo:  '☺',
	cor:      termbox.ColorRed,
	corFundo: termbox.ColorDefault,
	tangivel: true,
}
var personagem_2 = Elemento{
	simbolo:  '☺',
	cor:      termbox.ColorGreen,
	corFundo: termbox.ColorDefault,
	tangivel: true,
}
var personagem_3 = Elemento{
	simbolo:  '☺',
	cor:      termbox.ColorBlue,
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

// Servidor recebe (comando, jogador.posX, jogador.posY, elem, lastElem) e chama a função updatePos

var mapa [][]Elemento
var ultimoElementoSobPersonagem = vazio
var statusMsg string

var efeitoNeblina = false
var revelado [][]bool
var raioVisao int = 3

func inicializar(s *Servidor) {
	// Inicializa o mapa
	carregarMapa("mapa.txt")
	// Inicializa os jogadores
	s.Jogadores[0] = Jogador{ID: 0, Element: personagem_1, TX: 0, RX: 0, jogador.posX: -1, jogador.posY: -1, Online: false}
	s.Jogadores[1] = Jogador{ID: 1, Element: personagem_2, TX: 0, RX: 0, jogador.posX: -1, jogador.posY: -1, Online: false}
	s.Jogadores[2] = Jogador{ID: 2, Element: personagem_3, TX: 0, RX: 0, jogador.posX: -1, jogador.posY: -1, Online: false}
}
func (s *Servidor) ListenInput(j Jogador) { //TODO

}

func (s *Servidor) updateMap() { //TODO

}

func main() {

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
				// case personagem.simbolo:
				// 	// Atualiza a posição inicial do personagem
				// 	jogador.posX, jogador.posY = x, y
				// 	elementoAtual = vazio
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

func updatePos(novaPosX int, novaPosY int, elem Elemento, jogador Jogador) { // Cliente chama essa função para atualizar a posição do elemento

	if novaPosY >= 0 && novaPosY < len(mapa) && novaPosX >= 0 && novaPosX < len(mapa[novaPosY]) && mapa[novaPosY][novaPosX].tangivel == false {
		mapa[jogador.posY][jogador.posX] = ultimoElementoSobPersonagem        	// Restaura o elemento anterior
		ultimoElementoSobPersonagem = mapa[novaPosY][novaPosX] 					// Atualiza o elemento sob o personagem
		jogador.posX, jogador.posY = novaPosX, novaPosY                        	// Move o personagem
		mapa[jogador.posY][jogador.posX] = jogador.Element                      // Coloca o personagem na nova posição
	}
}

func interagir() { // TODO
	statusMsg = fmt.Sprintf("Interagindo em (%d, %d)", jogador.posX, jogador.posY)
}
