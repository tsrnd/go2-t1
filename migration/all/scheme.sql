CREATE DATABASE circle owner postgres encoding 'utf8';



/* Drop Tables */
DROP TABLE IF EXISTS user_app;


CREATE TABLE user_app
(
	id_user_app serial NOT NULL UNIQUE,
	user_name varchar(20),
    password varchar(50),
    phone varchar(11),
	user_profile_image_url varchar(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_app)
) WITHOUT OIDS;

ALTER SEQUENCE user_app_id_user_app_SEQ INCREMENT 1 RESTART 1;
