package main

import "fmt"

type No struct {
	Valor   int
	Proximo *No
}

type ListaCircular struct {
	Cabeca *No
}

func NovaListaCircular() *ListaCircular {
	return &ListaCircular{}
}

func (l *ListaCircular) InsereValorListaCircular(v int) {
	noParaInserir := &No{Valor: v}

	if l.Cabeca == nil {
		noParaInserir.Proximo = noParaInserir
		l.Cabeca = noParaInserir
	} else {
		noParaInserir.Proximo = l.Cabeca
		no := l.Cabeca
		for no.Proximo != l.Cabeca {
			no = no.Proximo
		}
		no.Proximo = noParaInserir
	}
}

func (l *ListaCircular) ExibeListaCircular() {
	if l.Cabeca == nil {
		fmt.Println("A lista está vazia")
		return
	}

	no := l.Cabeca
	for {
		fmt.Println(no.Valor)
		no = no.Proximo
		if no == l.Cabeca {
			break
		}
	}
	fmt.Println()
}

func (l *ListaCircular) RemoveValorListaCircular(v int) {
	if l.Cabeca == nil {
		return
	}

	no := l.Cabeca
	var noAnterior *No

	for {
		if no.Valor == v {
			if noAnterior == nil {
				noAnterior = no
				for noAnterior.Proximo != l.Cabeca {
					noAnterior = noAnterior.Proximo
				}
				l.Cabeca = no.Proximo
				noAnterior.Proximo = l.Cabeca
			} else {
				noAnterior.Proximo = no.Proximo
			}
			return
		}

		if no.Proximo == l.Cabeca {
			break
		}

		noAnterior = no
		no = no.Proximo
	}
}

func main() {
	lista := NovaListaCircular()
	lista.InsereValorListaCircular(1)
	lista.InsereValorListaCircular(2)
	lista.InsereValorListaCircular(3)
	lista.ExibeListaCircular()

	lista.RemoveValorListaCircular(2)
	lista.ExibeListaCircular()
}

