package model

import "errors"

//Imovel representa as informações de um imovel
type Imovel struct {
	X      int    `json:"coordenada_x"`
	Y      int    `json:"coordenada_y"`
	Nome   string `json:"nome"`
	Cidade string `json:"cidade"`
	valor  int
}

/*
ErrValorDeImovelInvalido é um erro que é apresentado quando é
atribuido a imóvel um valor muito baixo
*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido pra um imóvel")

/*
ErrValorDeImovelMuitoAlto é um erro que é apresentado quando é atribuido um valor
muito alto fora do mercado
*/
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return

}

//GetValor retorna o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
