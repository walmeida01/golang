package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/walmeida01/golang/avancado/json_unmarshall/model"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "http://jsonplaceholder.typicode.com/posts/2", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para a pagina. Erro", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina. Erro", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 201 {
		corpo, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da pagina. Erro", err.Error())
			return
		}
		fmt.Println("========================================= ")
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno JSON do servidor. Erro", err.Error())
			return
		}
		fmt.Printf("Post é: %+v\r\n", post)
	}

}
