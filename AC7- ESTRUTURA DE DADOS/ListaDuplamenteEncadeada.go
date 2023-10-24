package main

import "fmt"

type NoDuplo struct {
	Valor       int
	NóAnterior  *NoDuplo
	NóPróximo   *NoDuplo
}

type ListaDuplamenteEncadeada struct {
	PrimeiroNó *NoDuplo
	NóAnterior *NoDuplo
}

func NovaListaDuplamenteEncadeada() *ListaDuplamenteEncadeada {
	return &ListaDuplamenteEncadeada{}
}

func (l *ListaDuplamenteEncadeada) InsereValorListaOrdenada(v int) {
	novoNo := &NoDuplo{Valor: v}

	if l.PrimeiroNó == nil {
		l.PrimeiroNó = novoNo
		l.NóAnterior = novoNo
	} else if v <= l.PrimeiroNó.Valor { 
		novoNo.NóPróximo = l.PrimeiroNó
		l.PrimeiroNó.NóAnterior = novoNo
		l.PrimeiroNó = novoNo
	} else if v >= l.NóAnterior.Valor { 
		novoNo.NóAnterior = l.NóAnterior
		l.NóAnterior.NóPróximo = novoNo
		l.NóAnterior = novoNo
	} else { 
		atual := l.PrimeiroNó
		for atual != nil && v > atual.Valor {
			atual = atual.NóPróximo
		}

		novoNo.NóAnterior = atual.NóAnterior
		novoNo.NóPróximo = atual
		atual.NóAnterior.NóPróximo = novoNo
		atual.NóAnterior = novoNo
	}
}

func (l *ListaDuplamenteEncadeada) ExibeListaOrdenada() {
	if l.PrimeiroNó == nil {
		fmt.Println("A lista está vazia")
		return
	}

	no := l.PrimeiroNó
	for no != nil {
		fmt.Println(no.Valor)
		no = no.NóPróximo
	}
	fmt.Println()
}

func (l *ListaDuplamenteEncadeada) RemoveValorListaOrdenada(v int) {
	if l.PrimeiroNó == nil {  
		fmt.Println("A lista está vazia, não tem o que remover.")
		return
	}

	if l.PrimeiroNó.Valor == v {  
		l.PrimeiroNó = l.PrimeiroNó.NóPróximo
		if l.PrimeiroNó != nil {
			l.PrimeiroNó.NóAnterior = nil
		}
		return
	}

	if l.NóAnterior.Valor == v { 
		l.NóAnterior = l.NóAnterior.NóAnterior
		if l.NóAnterior != nil {
			l.NóAnterior.NóPróximo = nil
		}
		return
	}

	atual := l.PrimeiroNó
	for atual != nil && atual.Valor != v {
		atual = atual.NóPróximo
	}

	if atual != nil {
		atual.NóAnterior.NóPróximo = atual.NóPróximo
		atual.NóPróximo.NóAnterior = atual.NóAnterior
	}
}

func main() {
	lista := NovaListaDuplamenteEncadeada()

	lista.ExibeListaOrdenada()
	lista.InsereValorListaOrdenada(4)
	lista.ExibeListaOrdenada()
	lista.InsereValorListaOrdenada(2)
	lista.ExibeListaOrdenada()
	lista.InsereValorListaOrdenada(9)
	lista.ExibeListaOrdenada()
	lista.RemoveValorListaOrdenada(4) 
	lista.ExibeListaOrdenada()
	lista.RemoveValorListaOrdenada(99) 
	lista.ExibeListaOrdenada()
	lista.InsereValorListaOrdenada(5) 
	lista.ExibeListaOrdenada()
	lista.RemoveValorListaOrdenada(2) 
	lista.ExibeListaOrdenada()
	lista.RemoveValorListaOrdenada(9) 
	fmt.Println("Sucesso.")
}