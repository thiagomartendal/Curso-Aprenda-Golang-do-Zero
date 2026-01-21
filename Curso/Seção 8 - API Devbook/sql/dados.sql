insert into usuarios (nome, nick, email, senha) values
("Usuario 1", "us1", "us1@email.com", "$2a$10$OxzfgtzrvJmwEfuha5m55.JVVZbWBGNKe0TkIopdR/6qyJ/q1vViC"),
("Usuario 2", "us2", "us2@email.com", "$2a$10$OxzfgtzrvJmwEfuha5m55.JVVZbWBGNKe0TkIopdR/6qyJ/q1vViC"),
("Usuario 3", "us3", "us3@email.com", "$2a$10$OxzfgtzrvJmwEfuha5m55.JVVZbWBGNKe0TkIopdR/6qyJ/q1vViC"),
("Usuario 4", "us4", "us4@email.com", "$2a$10$OxzfgtzrvJmwEfuha5m55.JVVZbWBGNKe0TkIopdR/6qyJ/q1vViC");

insert into seguidores (usuario_id, seguidor_id) values (1, 2), (3, 1), (1, 3);

insert into publicacoes (titulo, conteudo, autor_id) values
("Apresentação 1", "Sou o usuário 1", 1),
("Apresentação 2", "Sou o usuário 2", 2),
("Apresentação 3", "Sou o usuário 3", 3);