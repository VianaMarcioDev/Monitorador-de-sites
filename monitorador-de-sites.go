package main

import "fmt"

func main(){
	nome := "Márcio"
	versao := 1.1
	fmt.Println("olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - inciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")

	var opcaoEscolhida int
	/*
	fmt.Scanf("%d", &opcaoEscolhida) Scanf precisa determinar o tipo
	*/
	fmt.Scan(&opcaoEscolhida)

	fmt.Println("O comando escolhido foi ", opcaoEscolhida)
}