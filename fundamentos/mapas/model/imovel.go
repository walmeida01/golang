package model

//Imovel representa as informações de um imovel
type Imovel struct {
	X      int    `json:"coordenada_x"`
	Y      int    `json:"coordenada_y"`
	Nome   string `json:"nome"`
	Cidade string `json:"cidade"`
	valor  int
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
