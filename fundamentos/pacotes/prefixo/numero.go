package prefixo

import "strconv"

//Capital representa o numero de telefone da capital de um estado/cidade
var Capital = 11

var teste = "xpto"

//TesteComPrefixo isto é um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
