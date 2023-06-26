package main

import (
	"fmt"
	"os"
)

func main() {

	boasVindas()
	menu()

	var opcaoEscolhida int
	/*
		fmt.Scanf("%d", &opcaoEscolhida) Scanf precisa determinar o tipo
	*/
	fmt.Scan(&opcaoEscolhida)

	switch opcaoEscolhida {
	case 0:
		fmt.Println("Saindo do programa.")
		os.Exit(0)
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Opção inválida!")
		os.Exit(-1)
	}

}

func boasVindas() {
	nome := "Márcio"
	versao := 1.1
	fmt.Println("Olá, Sr./Sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func menu() {
	fmt.Println("1 - inciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}
