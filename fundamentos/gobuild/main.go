package main

import (
	"encoding/json"
	"fmt"

	"github.com/walmeida01/golang/intermediario/gobuild/model"
)

/*
GOOS=windows GOARCH=386 go build -o meuapp.exe
GOOS=linux GOARCH=arm go build -o meuappraspberry -v github.com/walmeida01/golang/intermediario/gobuild
*/
func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	casa.SetValor(100)

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
