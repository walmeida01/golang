package main

import "fmt"

var (
	//Endereco é um valor importante e tem de ser publico
	Endereco string
	//Telefone é publica, telefone é privada
	Telefone            string  //endereco = ""
	quantidade, estoque int     //conforme o S.O, se for 32 int32 senao int64 // qtd = 0
	comprou             bool    //comprou = false
	valor               float64 //valor =0.00
	palavras            rune
)

func main() {
	teste := "Xpto"

	fmt.Printf("Endereco: %s\r\n", Endereco)
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("Comprou: %v\r\n", comprou)
	fmt.Printf("Palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
