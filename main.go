package main

import (
	"fmt"
)

type Conta struct {
	saldo float64
	limiteDiarioDeSaques int
	numeroDeSaques int
}

func main() {
	var conta Conta
	conta.saldo = 1000.00
	conta.limiteDiarioDeSaques = 2
	conta.numeroDeSaques = 0
	var opcao int

	fmt.Println("----- SISTEMA BANCARIO -----")
	fmt.Println("1-Saldo atual")
	fmt.Println("2-Depositar")
	fmt.Println("3-Sacar ")
	fmt.Println("4-Sair")

	fmt.Println("Digite a opcao desejada: ")
	fmt.Scanln(&opcao)

	if opcao == 4 {
		fmt.Println("OBRIGADO POR UTILIZAR O SISTEMA BANCARIO")
	}
}