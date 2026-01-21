package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

func Conectar() (*sql.DB, error) {
	conexao := "root:@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", conexao)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
