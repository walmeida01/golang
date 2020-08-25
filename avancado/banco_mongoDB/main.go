package main

import (
	"fmt"
	"net/http"

	"github.com/walmeida01/golang/avancado/banco_mongoDB/manipulador"
	"github.com/walmeida01/golang/avancado/banco_mongoDB/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	err = repo.AbreSessaoComMongo()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados: Erro ao abrir a sessao com o MongoDB ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Escutando....")
	http.ListenAndServe(":8181", nil)

}
