package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Importação implícita
)

func main() {
	// conexao := "usuario:senha@/banco"
	conexao := "root:@/devbook?charset=utf8&parseTime=True&loc=Local"
	// charset: conjunto de caracteres - parseTime: formatação de datas - loc: localização
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
