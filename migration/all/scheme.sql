CREATE DATABASE circle owner postgres encoding 'utf8';



/* Drop Tables */
DROP TABLE IF EXISTS users;


CREATE TABLE users
(
	id serial NOT NULL UNIQUE,
	username varchar(36),
    password varchar(20),
    phone varchar(11),
	avatar varchar(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_app)
) WITHOUT OIDS;

ALTER SEQUENCE id_users_SEQ INCREMENT 1 RESTART 1;
