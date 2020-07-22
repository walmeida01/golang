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

	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{10, 20, "Chacara nova", 20000}
	fmt.Printf("Casa é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("chacara é: %p - %+v\r\n", &chacara, chacara)
}

//*Imovel ponteiro do valor
func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5

}
