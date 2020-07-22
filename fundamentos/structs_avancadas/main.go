package main

import (
	"encoding/json"
	"fmt"

	"github.com/walmeida01/golang/fundamentos/structs_avancadas/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(6000)

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O Valor da casa é: %+v\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))

}
