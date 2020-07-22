package operadora

import (
	"strconv"

	"github.com/walmeida01/golang/fundamentos/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
