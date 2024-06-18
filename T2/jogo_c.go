package main

import (
    "fmt"
    "net/rpc"
    "os"

	"github.com/nsf/termbox-go"
)
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


func main() {

    if len(os.Args) != 2 {
        fmt.Println("Uso:", os.Args[0], " <maquina>")
        return
    }

    porta := 8973
    maquina := os.Args[1]

    client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        return
    }

    var nota float64

    err = client.Call("Servidor.ObtemNota", nome, &nota)
    if err != nil {
        fmt.Println("Erro ao obter nota:", err)
    } else {
        fmt.Printf("Nome: %s\n", nome)
        fmt.Printf("Nota: %.2f\n", nota)
    }
}
