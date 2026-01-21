package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	stmt, err := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.IdAutor)
	if err != nil {
		return 0, err
	}

	idInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(idInserido), nil
}

func (repositorio Publicacoes) BucarPorID(idPublicacao uint64) (modelos.Publicacao, error) {
	linha, err := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u on
		u.id = p.autor_id where p.id = ?
	`, idPublicacao)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if err := linha.Scan(
			&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo,
			&publicacao.IdAutor, &publicacao.Curtidas, &publicacao.CriadaEm,
			&publicacao.NickAutor); err != nil {
			return modelos.Publicacao{}, err
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Bucar(idUsuario uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
		select distinct p.*, u.nick from publicacoes p inner join usuarios u on
		u.id = p.autor_id inner join seguidores s on
		p.autor_id = s.usuario_id where u.id = ? or s.seguidor_id = ?
		order by 1 desc
	`, idUsuario, idUsuario)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err := linhas.Scan(
			&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo,
			&publicacao.IdAutor, &publicacao.Curtidas, &publicacao.CriadaEm,
			&publicacao.NickAutor); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(idPublicacao uint64, publicacao modelos.Publicacao) error {
	stmt, err := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, idPublicacao); err != nil {
		return err
	}

	return nil
}

func (repositorio Publicacoes) Deletar(idPublicacao uint64) error {
	stmt, err := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(idPublicacao); err != nil {
		return err
	}

	return nil
}

func (repositorio Publicacoes) BuscarPorUsuario(idUsuario uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p join usuarios u on
		u.id = p.autor_id where p.autor_id = ?
	`, idUsuario)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err := linhas.Scan(
			&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo,
			&publicacao.IdAutor, &publicacao.Curtidas, &publicacao.CriadaEm,
			&publicacao.NickAutor); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Curtir(idPublicacao uint64) error {
	stmt, err := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(idPublicacao); err != nil {
		return err
	}

	return nil
}

func (repositorio Publicacoes) Descurtir(idPublicacao uint64) error {
	stmt, err := repositorio.db.Prepare(`
		update publicacoes set curtidas = case when curtidas > 0 then
		curtidas - 1 else curtidas end where id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(idPublicacao); err != nil {
		return err
	}

	return nil
}
