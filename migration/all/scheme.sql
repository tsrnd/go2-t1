CREATE DATABASE circle owner postgres encoding 'utf8';



/* Drop Tables */
DROP TABLE IF EXISTS users;


CREATE TABLE users
(
	id serial NOT NULL UNIQUE,
	username char(36) NOT NULL UNIQUE,
	password varchar(20),
	phone varchar(11)
	avatar varchar(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE id_users_SEQ INCREMENT 1 RESTART 1;