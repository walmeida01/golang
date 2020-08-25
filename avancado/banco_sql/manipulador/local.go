package manipulador

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/walmeida01/golang/avancado/banco_sql/repo"

	"github.com/walmeida01/golang/avancado/banco_sql/model"
)

//Local é o manipulador da requisição da rota /local/
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido, Verifique.", http.StatusBadRequest)
		fmt.Println("[local] Erro ao converter o número enviado", err.Error())
	}

	sql := "select country, city, telcode from place where telcode = ?"
	rows, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse número.", http.StatusInternalServerError)
		fmt.Println("[local] Nao foi possivel executar a query", sql, "Erro: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possivel pesquisar esse número.", http.StatusBadRequest)
			fmt.Println("[local] Nao foi possivel fazer o binding dos dados do banco na struct", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[local] Erro na execução do modelo", err.Error())
	}

	sql = "insert into logquery (daterequest) values (?)"
	result, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[local] Erro na inclusao do log: ", sql, " - ", err.Error())
	}

	rowsAffctd, err := result.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o numero de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}

	fmt.Println("Sucesso! ", rowsAffctd, " linha(s) afetada(s)")
}
