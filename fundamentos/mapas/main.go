package main

import (
	"fmt"

	"github.com/walmeida01/golang/fundamentos/mapas/model"
)

var cache map[string]model.Imovel

func main() {
	//mapas é uma coleção de dados, tem indices(chaves)

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Azul"
	casa.Cidade = "RJ"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(6000)

	cache["Casa Azul"] = casa

	fmt.Println("Há uma casa azul no cache ?")
	imovel, achou := cache["Casa Azul"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.Cidade = "SP"
	apto.X = 20
	apto.Y = 40
	apto.SetValor(7500)

	//para inserir pode usar qualquer tipo
	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? :", len(cache))

	//primeiro obj é a chave e o segundo é o próprio obj
	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}

	imovel, achou = cache["Apto Azul"]
	if achou {
		delete(cache, imovel.Nome)
	}
	fmt.Println("Quantos itens há no cache agora? :", len(cache))
}
