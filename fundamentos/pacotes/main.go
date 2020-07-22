package main

import (
	"fmt"

	"github.com/walmeida01/golang/fundamentos/pacotes/operadora"
	"github.com/walmeida01/golang/fundamentos/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do S.O
var NomeDoUsuario = "Will"

func main() {
	fmt.Printf("Nome do usuário é: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("O valor de teste é: %s\r\n", prefixo.TesteComPrefixo)
}
