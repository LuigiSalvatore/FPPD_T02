package main

//	O servidor só precisa gerar o mapa
//	Deve ser implementadas funções que modifiquem o mapa, movem os jogadores, e interajam com os elementos
//
//
//
import (
	"bufio"
	"fmt"
	"net"
	"net/rpc"
	"os"

	"github.com/nsf/termbox-go"
)

// Define os elementos do jogo
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
type Servidor struct {
	Jogadores                   [3]Jogador
	mapa                        [][]Elemento
	ultimoElementoSobPersonagem Elemento
	statusMsg                   string
	efeitoNeblina               bool
	revelado                    [][]bool
	raioVisao                   int
	mapa_Inicializado           bool
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
	Cor:      termbox.ColorBlack | termbox.AttrBold | termbox.AttrDim,
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

// Servidor recebe (comando, jogador.PosX, jogador.PosY, elem, lastElem) e chama a função updatePos

// var mapa [][]Elemento
// var ultimoElementoSobPersonagem = vazio
// var statusMsg string
// var efeitoNeblina = false
// var revelado [][]bool
// var raioVisao int = 3
var meuCu int = 0

func (s *Servidor) inicializar() {
	s.ultimoElementoSobPersonagem = vazio
	s.efeitoNeblina = false
	s.raioVisao = 3
	s.mapa_Inicializado = false
	// Inicializa os jogadores
	s.Jogadores[0] = Jogador{ID: 0, Element: personagem_1, PosX: 10, PosY: 3, Online: false}
	s.Jogadores[1] = Jogador{ID: 1, Element: personagem_2, PosX: 11, PosY: 3, Online: false}
	s.Jogadores[2] = Jogador{ID: 2, Element: personagem_3, PosX: 12, PosY: 3, Online: false}
	// Inicializa o mapa
	s.carregarMapa("mapa.txt")
}
func (s *Servidor) SendMapa(id string, clientMap *[][]Elemento) error { // cliente manda seu mapa, servidor Carrega o mapa
	if s.mapa_Inicializado {
		*clientMap = s.mapa
		// fmt.Println("Mapa enviado para", id)
		return nil
	}
	return fmt.Errorf("Mapa ainda não Inicializado")
}

func (s *Servidor) ListenInput(ev rune, id_c *int) error { /*TODO*/
	if *id_c >= 0 || *id_c < 2 {

		fmt.Println("Jogador", s.Jogadores[*id_c].ID, "recebeu", ev, "ev =", string(ev))
		fmt.Println("Jogador", s.Jogadores[*id_c].ID, "PosX:", s.Jogadores[*id_c].PosX, "PosY:", s.Jogadores[*id_c].PosY)
		switch ev {
		case 'w':
			go s.updatePos(s.Jogadores[*id_c].PosX, s.Jogadores[*id_c].PosY-1, s.Jogadores[*id_c].Element, &s.Jogadores[*id_c])
		case 'a':
			go s.updatePos(s.Jogadores[*id_c].PosX-1, s.Jogadores[*id_c].PosY, s.Jogadores[*id_c].Element, &s.Jogadores[*id_c])
		case 's':
			go s.updatePos(s.Jogadores[*id_c].PosX, s.Jogadores[*id_c].PosY+1, s.Jogadores[*id_c].Element, &s.Jogadores[*id_c])
		case 'd':
			go s.updatePos(s.Jogadores[*id_c].PosX+1, s.Jogadores[*id_c].PosY, s.Jogadores[*id_c].Element, &s.Jogadores[*id_c])
		case 'e':
			go s.interact(ev, *id_c)
		}
		fmt.Println("Jogador", s.Jogadores[*id_c].ID, "PosX:", s.Jogadores[*id_c].PosX, "PosY:", s.Jogadores[*id_c].PosY)
		return nil
	}
	return fmt.Errorf("Jogador não encontrado")
}

func (s *Servidor) interact(ev rune, id_c int) { /*idk what TODO*/
	fmt.Println("Interagindo com", ev)
}
func (s *Servidor) GetID(trash string, id *int) error { /*DONE*/

	id = &meuCu
	meuCu++
	// for i := 0; i < 3; i++ {
	// 	if !s.Jogadores[i].Online {
	// 		s.Jogadores[i].Online = true
	// 		*id = s.Jogadores[i].ID
	// 		fmt.Println("Jogador", i, "conectado")
	// 		fmt.Println("Enviado:", s.Jogadores[i].ID, s.Jogadores[i].PosX, s.Jogadores[i].PosY, s.Jogadores[i].Element.Simbolo, s.Jogadores[i].Element.Cor, s.Jogadores[i].Element.CorFundo, s.Jogadores[i].Element.Tangivel)
	// 		fmt.Println("Copiado: Jogador", *id, "PosX:", s.Jogadores[i].PosX, "PosY:", s.Jogadores[i].PosY, "Elemento:", s.Jogadores[i].Element.Simbolo, s.Jogadores[i].Element.Cor, s.Jogadores[i].Element.CorFundo, s.Jogadores[i].Element.Tangivel)
	// 		return nil
	// 	}
	// }
	// return fmt.Errorf("Não há mais jogadores disponíveis")
	return nil
}

func (s *Servidor) AckPlayer(trash string, j *Jogador) error {
	fmt.Println("Confirmando jogador", j.ID)
	fmt.Println("Jogador", j.ID, "PosX:", j.PosX, "PosY:", j.PosY, "Elemento:", j.Element.Simbolo, j.Element.Cor, j.Element.CorFundo, j.Element.Tangivel)
	for i := 0; i < 3; i++ {
		fmt.Println("Comparando", s.Jogadores[i].ID, j.ID, s.Jogadores[i].PosX, j.PosX, s.Jogadores[i].PosY, j.PosY)
		if s.Jogadores[i].ID == j.ID && s.Jogadores[i].PosX == j.PosX && s.Jogadores[i].PosY == j.PosY {
			fmt.Println("Jogador", i, "confirmado")
			return nil
		}
	}
	fmt.Println("Jogador não encontrado")
	return fmt.Errorf("Jogador não encontrado")
}
func main() {
	porta := 8973
	servidor := new(Servidor)
	servidor.inicializar()
	if servidor.mapa == nil || len(servidor.mapa) == 0 {
		fmt.Println("Mapa não inicializado")
		return
	} else {
		fmt.Println("Mapa inicializado")
	}
	rpc.Register(servidor)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", porta))
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}

	fmt.Println("Servidor aguardando conexões na porta", porta)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		} else {
			fmt.Println("Conexão aceita:", conn.RemoteAddr())
		}
		go rpc.ServeConn(conn)
	}

}

func (s *Servidor) carregarMapa(nomeArquivo string) {
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
		for x, char := range linhaTexto {
			elementoAtual := vazio
			switch char {
			case parede.Simbolo:
				elementoAtual = parede
			case barreira.Simbolo:
				elementoAtual = barreira
			case vegetacao.Simbolo:
				elementoAtual = vegetacao
			}
			linhaElementos = append(linhaElementos, elementoAtual)
			x++
		}
		s.mapa = append(s.mapa, linhaElementos)
		y++
	}
	// Coloca o personagem na Posição inicial
	// s.mapa[s.Jogadores[0].PosY][s.Jogadores[0].PosX] = s.Jogadores[0].Element
	// s.mapa[s.Jogadores[1].PosY][s.Jogadores[1].PosX] = s.Jogadores[1].Element
	// s.mapa[s.Jogadores[2].PosY][s.Jogadores[2].PosX] = s.Jogadores[2].Element

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	s.mapa_Inicializado = true
}

func (s *Servidor) updatePos(novaPosX int, novaPosY int, elem Elemento, j *Jogador) { // Cliente chama essa função para atualizar a Posição do elemento
	fmt.Println("Jogador com ID = ", j.ID, "updatedPos", novaPosX, novaPosY)
	if novaPosY >= 0 && novaPosY < len(s.mapa) && novaPosX >= 0 && novaPosX < len(s.mapa[novaPosY]) && s.mapa[novaPosY][novaPosX].Tangivel == false {
		s.mapa[j.PosY][j.PosX] = s.ultimoElementoSobPersonagem     // Restaura o elemento anterior
		s.ultimoElementoSobPersonagem = s.mapa[novaPosY][novaPosX] // Atualiza o elemento sob o personagem
		j.PosX, j.PosY = novaPosX, novaPosY                        // Move o personagem
		s.mapa[j.PosY][j.PosX] = j.Element                         // Coloca o personagem na nova Posição
	}
	fmt.Println("Tentei mover para", novaPosX, novaPosY)
}
