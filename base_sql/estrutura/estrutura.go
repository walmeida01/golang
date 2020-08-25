package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("[estrutura] Erro ao executar comando: ", sql, ", no banco:", err.Error())
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		fmt.Println("[main] Erro ao conectar ao banco", err.Error())
	}
	defer db.Close()

	exec(db, "create database if not exists golang")
}
