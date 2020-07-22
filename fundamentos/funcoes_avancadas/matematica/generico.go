package matematica

//Calculo executa qualquer tipo de calculo, basta enviar a funcao desejada
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica x vezes o y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor da funcao
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto retorna o resultado e o resto da funcao
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
