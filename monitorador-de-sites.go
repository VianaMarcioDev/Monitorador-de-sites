package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	boasVindas()
	menu()	
	opcaoEscolhida := leComando()

	switch opcaoEscolhida {
	case 0:
		fmt.Println("Saindo do programa.")
		os.Exit(0)
	case 1:
		iniciarMonitoramento()
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
func leComando() int{
	var opcaoEscolhida int 
	fmt.Scan(&opcaoEscolhida)
	fmt.Println("O comando escolhido foi", opcaoEscolhida)
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	site := "http://www.alura.com.br/"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200{
		fmt.Println("Site", site,"foi carregado com sucesso!")
	}else{
		fmt.Println("Site com problemas! Staus do problema", resp.StatusCode)
	}
}