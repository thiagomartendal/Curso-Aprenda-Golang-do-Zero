package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	stmt, err := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoID, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoID), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // Formato: %valor%

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick,
		nomeOuNick,
	)
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorId(id uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", id)
	if err != nil {
		return modelos.Usuario{}, nil
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	stmt, err := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) Deletar(ID uint64) error {
	stmt, err := repositorio.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(idUsuario, idSeguidor uint64) error {
	// ignore impede a inserção de dados que já estão no banco
	stmt, err := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(idUsuario, idSeguidor); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguir(idUsuario, idSeguidor uint64) error {
	stmt, err := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(idUsuario, idSeguidor); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(idUsuario uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join
		seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, idUsuario)

	if err != nil {
		return nil, err
	}

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if err = linhas.Scan(
			&usuario.ID, &usuario.Nome, &usuario.Nick,
			&usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSeguindo(idUsuario uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join
		seguidores s on u.id = s.usuario_id where s.seguidor_id = ?
	`, idUsuario)

	if err != nil {
		return nil, err
	}

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if err = linhas.Scan(
			&usuario.ID, &usuario.Nome, &usuario.Nick,
			&usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSenha(idUsuario uint64) (string, error) {
	linha, err := repositorio.db.Query("select senha from usuarios where id = ?", idUsuario)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if err := linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}

	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(idUsuario uint64, senha string) error {
	stmt, err := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(senha, idUsuario); err != nil {
		return err
	}

	return nil
}
