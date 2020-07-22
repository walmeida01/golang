package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3
	fmt.Print("O número ", numero, " se escreve assim:")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")

	}

	fmt.Println("Você é da familia do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Deixa pra lá né rs")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ok, voce pode dormir até mais tarde")
	default:
		fmt.Println("Vamos lá que hoje é dia de trabalho")
	}

	numero = 9
	fmt.Println("Este numero cabe num digito?")
	switch {
	case numero < 10:
		fmt.Println("SIM")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois digitos...")
	case numero >= 100:
		fmt.Println("Não dá o numero é muito grande")
	}
}
