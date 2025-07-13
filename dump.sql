CREATE TABLE public.users (
	uid varchar(26) NOT null,
	handle varchar(30) NOT NULL,
	email varchar(75),
	phone_number varchar(100),
	passkey varchar(100) NOT NULL,
	CONSTRAINT uids_handle_un UNIQUE (handle),
	CONSTRAINT uids_pk PRIMARY KEY (uid)
);