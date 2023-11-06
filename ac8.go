package main

import "fmt"

type no struct {
	valor    int
	esquerda *no
	direita  *no
}

type arvoreBinaria struct {
	raiz *no
}

func buscarValor(no *no, valor int) bool {
	if no == nil {
		return false
	}
	if valor == no.valor {
		return true
	}
	if valor < no.valor {
		return buscarValor(no.esquerda, valor)
	}
	return buscarValor(no.direita, valor)
}

func inserir(arvore *arvoreBinaria, valor int) string {
	novoNo := &no{valor: valor}
	if arvore.raiz == nil {
		arvore.raiz = novoNo
		return "Valor inserido como raiz"
	}
	noAtual := arvore.raiz
	for {
		if valor == noAtual.valor {
			return "Valor já existe na árvore"
		}
		if valor < noAtual.valor {
			if noAtual.esquerda == nil {
				noAtual.esquerda = novoNo
				break
			}
			noAtual = noAtual.esquerda
		} else {
			if noAtual.direita == nil {
				noAtual.direita = novoNo
				break
			}
			noAtual = noAtual.direita
		}
	}
	return "Valor inserido na árvore"
}

func main() {
	arvore := arvoreBinaria{}

	fmt.Println(inserir(&arvore, 10))
	fmt.Println(inserir(&arvore, 5))
	fmt.Println(inserir(&arvore, 15))
	fmt.Println(inserir(&arvore, 15))

	valorBusca := 15
	if buscarValor(arvore.raiz, valorBusca) {
		fmt.Printf("Valor %d encontrado na árvore\n", valorBusca)
	} else {
		fmt.Printf("Valor %d não encontrado na árvore\n", valorBusca)
	}
}

