package main

import (
	"fmt"

	"github.com/walmeida01/golang/fundamentos/funcoes_avancadas/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado do multiplicador é: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma é: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da divisão foi: %d e o resto da divisão foi %d \r\n", resultado, resto)
}
