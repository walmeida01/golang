package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/walmeida01/golang/avancado/web_post/model"
)

//http://requestbin.net/r/sed0xfse?inspect

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	user := model.Usuario{}
	user.ID = 1
	user.Nome = "Will2"

	contentSend, err := json.Marshal(user)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o objeto JSON para o usuario. Erro", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://requestbin.net/r/sed0xfse", bytes.NewBuffer(contentSend))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para a pagina. Erro", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço do request bin. Erro", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 201 {
		corpo, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao chamar o serviço do request bin. Erro", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}
