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
	posX    int
	posY    int
	Online  bool
}
type Servidor struct {
	Jogadores 					[3]Jogador
	mapa 						[][]Elemento
	ultimoElementoSobPersonagem Elemento
	statusMsg 					string
	efeitoNeblina  				bool
	revelado 					[][]bool
	raioVisao 					int 
	mapa_Inicializado 			bool

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

// var mapa [][]Elemento
// var ultimoElementoSobPersonagem = vazio
// var statusMsg string
// var efeitoNeblina = false
// var revelado [][]bool
// var raioVisao int = 3

func inicializar(s *Servidor) {
	ultimoElementoSobPersonagem = vazio
	efeitoNeblina = false
	raioVisao = 3
	mapa_Inicializado = false
	// Inicializa o mapa
	carregarMapa("mapa.txt")
	// Inicializa os jogadores
	s.Jogadores[0] = Jogador{ID: 0, Element: personagem_1, TX: 0, RX: 0, posX: -1, posY: -1, Online: false}
	s.Jogadores[1] = Jogador{ID: 1, Element: personagem_2, TX: 0, RX: 0, posX: -1, posY: -1, Online: false}
	s.Jogadores[2] = Jogador{ID: 2, Element: personagem_3, TX: 0, RX: 0, posX: -1, posY: -1, Online: false}
}
func (s *Servidor) sendMapa(clientMap *[][]Elemento) { // cliente manda seu mapa, servidor Carrega o mapa
	if(mapa_Inicializado){
		*clientMap = mapa
		return nil
	}
	return fmt.Errorf("Mapa ainda não Inicializado")
}
func (s *Servidor) listenInput(j Jogador) { //TODO

}

func (s *Servidor) updateMap(j Jogador) { //TODO

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
	mapa_Inicializado = true
}

func updatePos(novaPosX int, novaPosY int, elem Elemento, j Jogador) { // Cliente chama essa função para atualizar a posição do elemento

	if novaPosY >= 0 && novaPosY < len(mapa) && novaPosX >= 0 && novaPosX < len(mapa[novaPosY]) && mapa[novaPosY][novaPosX].tangivel == false {
		mapa[j.posY][j.posX] = ultimoElementoSobPersonagem        	// Restaura o elemento anterior
		ultimoElementoSobPersonagem = mapa[novaPosY][novaPosX] 					// Atualiza o elemento sob o personagem
		j.posX, j.posY = novaPosX, novaPosY                        	// Move o personagem
		mapa[j.posY][j.posX] = j.Element                      // Coloca o personagem na nova posição
	}
}

func interagir() { // TODO
	statusMsg = fmt.Sprintf("Interagindo em (%d, %d)", jogador.posX, jogador.posY)
}
