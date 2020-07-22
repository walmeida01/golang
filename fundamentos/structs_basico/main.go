package main

import "fmt"

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu Apt", 750000}
	fmt.Printf("O Apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     10,
		Nome:  "Chacara",
		X:     2,
		valor: 50,
	}

	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Lar doce Lar"
	casa.valor = 450000
	casa.X = 10
	casa.Y = 20

	fmt.Printf("A casa é: %+v\r\n", casa)

}
