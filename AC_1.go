package main

import (
	"fmt"
)

func e_primo(numero int) (bool, []int) {

	divisores := []int{1}
	Primo := true
	if numero <= 1 {
		return false, []int{}
	}

	for i := 2; i <= numero/2; i++ {
		if numero%i == 0 {
			divisores = append(divisores, i)
			Primo = false
		}
	}

	if Primo {
		return true, []int{}
	}

	return Primo, divisores
}

func fibo(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	prev1 := 1
	prev2 := 1
	resultado := 0

	for i := 2; i < n; i++ {
		resultado = prev1 + prev2
		prev1, prev2 = prev2, resultado
	}

	return resultado
}

func diaDaSemana(numero int) string {
	if numero == 1 {
		return "Domingo"
	} else if numero == 2 {
		return "Segunda-feira"
	} else if numero == 3 {
		return "Terça-feira"
	} else if numero == 4 {
		return "Quarta-feira"
	} else if numero == 5 {
		return "Quinta-feira"
	} else if numero == 6 {
		return "Sexta-feira"
	} else if numero == 7 {
		return "Sábado"
	} else {
		return "Valor inválido"
	}
}
func e_bissexto(ano int) bool {
	if ano%4 == 0 && ano%100 != 0 {
		return true
	}
	return false
}
func main() {
	fmt.Println(e_primo(37))
	fmt.Println(fibo(30))
	fmt.Println(diaDaSemana(5))
	fmt.Println(e_bissexto(2024))
}
