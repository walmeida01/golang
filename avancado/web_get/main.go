package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second + 30,
	}

	response, err := client.Get("https://google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina. Erro", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da pagina. Erro", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para a pagina. Erro", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina. Erro", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da pagina. Erro", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}
