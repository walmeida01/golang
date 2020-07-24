package main

import (
	"fmt"

	"github.com/walmeida01/golang/fundamentos/interfaces/model"
)

func main() {

	wanda := model.Ave{}
	wanda.Nome = "Wanda da silva"

	queroAcordarComUmCacarejo(wanda)
	queroOuvirUmaPataNoLago(wanda)
}

//QueroAcordarComUmCacarejo representa
func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
