package main

import (
	"fmt"
	"calculadora/funcoesCalculadora"
)

func main() {
	var {
		opcaoEscolha int
		num_1 num_2 int
	}
	
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scanf(num_1)
	fmt.Print("Digite o Segundo valor: ")
	fmt.Scanf(num_2)

	fmt.Print("Escolha uma opção dentre [1] - Somar, [2] - Subtrair, [3] - Multiplicar ou [4] - dividir\n > ")
	fmt.Scanf(opcaoEscolha)

	switch opcaoEscolha {
	case 1:
		fmt.Printf("%d+%d=%d", num_1, num_2, funcoesCalculadora.Somar(num_1, num_2))
	case 2:
		fmt.Printf("%d-%d=%d", num_1, num_2, funcoesCalculadora.Subtrair(num_1, num_2))
	case 3:
		fmt.Printf("%d*%d=%d", num_1, num_2, funcoesCalculadora.Multiplicar(num_1, num_2))
	case 4:
		fmt.Printf("%d/%d=%d", num_1, num_2, funcoesCalculadora.Dividir(num_1, num_2))
	case default:
		fmt.Print("Opção inválida!")
	}
}