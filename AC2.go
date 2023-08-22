package main

import "fmt"

type Contato struct {
    Nome          string
    Email         string
    ListadeContato []string
}

func (c *Contato) DefinirEmail(novoEmail string) {
    c.Email = novoEmail
}

func (c *Contato) AdicionarContato(telefone string) {
    for i := range c.ListadeContato {
        if c.ListadeContato[i] == "" {
            c.ListadeContato[i] = telefone
            return
        }
    }
    c.ListadeContato = append(c.ListadeContato, telefone)
}

func (c *Contato) ExcluirContato() {
    for i := len(c.ListadeContato) - 1; i >= 0; i-- {
        if c.ListadeContato[i] != "" {
            c.ListadeContato[i] = ""
            return
        }
    }
}

func main() {
    contato1 := Contato{
        Nome:          "Pedro",
        Email:         "peedrplustosa",
        ListadeContato: make([]string, 5), // Inicializa com 5 elementos vazios
    }

    fmt.Println(contato1.Nome, contato1.Email, contato1.ListadeContato)

    contato1.DefinirEmail("pedroplustosa@gmail.com")

    fmt.Println(contato1.Nome, contato1.Email)

    contato1.AdicionarContato("1234-5678")

    fmt.Println(contato1.Nome, contato1.Email, contato1.ListadeContato)
    
    contato1.ExcluirContato()
    
    fmt.Println(contato1.Nome, contato1.Email, contato1.ListadeContato)
}
