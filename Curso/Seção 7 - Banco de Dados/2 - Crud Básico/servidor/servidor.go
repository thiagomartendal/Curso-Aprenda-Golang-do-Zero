package servidor

import (
	"crud/banco"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func conexaoBD(w http.ResponseWriter) *sql.DB {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao se conectar ao banco de dados."))
		os.Exit(1)
	}
	return db
}

func obterID(url string, w http.ResponseWriter) uint64 {
	vars := strings.Split(url, "/")
	ID, err := strconv.ParseUint(vars[2], 10, 64)
	if err != nil {
		w.Write([]byte("Erro durante a conversão de tipo do identificador."))
		os.Exit(1)
	}
	return ID
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler a requisição."))
		return
	}

	var usr usuario

	if err := json.Unmarshal(req, &usr); err != nil {
		w.Write([]byte("Erro na conversão de dados do usuário."))
	}

	db := conexaoBD(w)
	defer db.Close()

	// Prepare Statement - Define o formato da consulta sql
	stmt, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro na definição da consulta."))
		return
	}
	defer stmt.Close()

	insercao, err := stmt.Exec(usr.Nome, usr.Email)
	if err != nil {
		w.Write([]byte("Erro ao inserir os dados no banco."))
		return
	}

	id, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o identificador do usuário inserido."))
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, "Usuário %d inserido com sucesso.", id) // Escreve no ResponseWriter w
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db := conexaoBD(w)
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários."))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário."))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON."))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	ID := obterID(r.URL.Path, w)

	db := conexaoBD(w)
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		fmt.Fprintf(w, "Erro ao buscar o usuário com o identificador %d.", ID)
		return
	}

	var usuario usuario
	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário."))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		fmt.Fprintf(w, "Erro ao converter o usuário com o identificador %d para JSON.", ID)
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	ID := obterID(r.URL.Path, w)

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição."))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter as informações do usuário."))
		return
	}

	db := conexaoBD(w)
	defer db.Close()

	stmt, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro na definição da consulta."))
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		fmt.Fprintf(w, "Erro ao atualizar o usuário com identificador %d.", ID)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	ID := obterID(r.URL.Path, w)

	db := conexaoBD(w)
	defer db.Close()

	stmt, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro na definição da consulta."))
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(ID); err != nil {
		fmt.Fprintf(w, "Erro ao deletar o usuário com identificador %d.", ID)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
