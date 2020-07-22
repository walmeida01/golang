package main

import (
	"fmt"

	"github.com/walmeida01/golang/fundamentos/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado do multiplicador é: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma é: %d\r\n", resultado)
}
