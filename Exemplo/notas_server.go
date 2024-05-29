package main

import (
    "fmt"
    "net"
    "net/rpc"
)

// Estrutura para representar um aluno
type Aluno struct {
    Nome string
    Nota float64
}

// Estrutura para o servidor
type Servidor struct {
    alunos []Aluno
}

// Método para inicializar a lista de alunos no servidor
func (s *Servidor) inicializar() {
    s.alunos = []Aluno{
		{"Alexandre", 9.5},
		{"Barbara",   8.5},
		{"Joao",      6.5},
		{"Maria",     9.0},
		{"Paulo",    10.0},
		{"Pedro",     7.0},
	}
}

// Método remoto que retorna a nota de um aluno dado o seu nome
func (s *Servidor) ObtemNota(nome string, nota *float64) error {
    for _, aluno := range s.alunos {
        if aluno.Nome == nome {
            *nota = aluno.Nota
            return nil
        }
    }
    return fmt.Errorf("Aluno %s não encontrado", nome)
}

func main() {
    porta := 8973
    servidor := new(Servidor)
    servidor.inicializar()

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
        }
        go rpc.ServeConn(conn)
    }
}
