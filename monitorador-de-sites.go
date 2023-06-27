package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	for {
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

}

func boasVindas() {
	nome := "Márcio"
	versao := 1.1
	fmt.Println("Olá, Sr./Sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func menu() {
	fmt.Println("1 - iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}
func leComando() int {
	var opcaoEscolhida int
	fmt.Scan(&opcaoEscolhida)
	fmt.Println("O comando escolhido foi", opcaoEscolhida)
	return opcaoEscolhida
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := [] string {"http://www.alura.com.br/", "https://www.caelum.com.br", "https://random-status-code.herokuapp.com"}
	for i := 0; i < monitoramentos; i++{
		for i, site := range sites{
			fmt.Println("Testando site", (i + 1),":", site)
			testaSite(site)
			fmt.Println("")
		}
			time.Sleep(delay * time.Second)
			fmt.Println("")
	}
	
	fmt.Println("")
	
}
func testaSite(site string){
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site com problemas! Staus do problema", resp.StatusCode)
	}
}
