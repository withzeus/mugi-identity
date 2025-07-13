CREATE SEQUENCE uid_seq;
CREATE TABLE uids (
	uid int NOT null DEFAULT nextval('uid_seq'),
	handle varchar(30) NOT NULL,
	email varchar(75),
	phone_number varchar(100),
	passkey varchar(100) NOT NULL,
	CONSTRAINT uids_handle_un UNIQUE (handle),
	CONSTRAINT uids_pk PRIMARY KEY (uid)
);