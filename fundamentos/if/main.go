package main

import "fmt"

func main() {

	meses := 0
	situacao := true
	cidade := "SP"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Usuário esta devendo")
	}

	if !situacao {
		fmt.Println("Usuário está adimplente")
	}

	if cidade == "SP" {
		fmt.Println("O cliente vive no estado de SP")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente ? ", descricao)
		return
	}

	//variaveis descricao, status, são validas somente no if, fora disso, não compilam.
	fmt.Println("Obrigado por nos consultar!")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo"
		return
	}
	descricao = "O cliente está em dia."
	return

}
