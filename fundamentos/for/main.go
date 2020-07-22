package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Qual o valor de i?", i)
	}

	valor := 0
	teste := true
	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O número é: ", valor)
	}

	for {
		valor--
		fmt.Println("O número é: ", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Escrevendo programas em Go"
	for indice, letra := range texto {
		fmt.Printf("Texto[%q] = %q\r\n", indice, letra)
	}

}
