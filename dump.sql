DROP TABLE IF EXISTS public.clients;

CREATE TABLE public.clients (
	id varchar(26) NOT null,
	name varchar(30) NOT NULL,
	secret varchar(75),
	website varchar(50) NOT NULL,
	logo varchar(50),
	redirect_uri varchar(100) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deeated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    upleted_at TIMESTAMP WITH TIME ZONE,
	CONSTRAINT clients_id PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users (
	uid varchar(26) NOT null,
	handle varchar(30) NOT NULL,
	email varchar(75),
	phone_number varchar(100),
	passkey varchar(100) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
	CONSTRAINT users_handle_un UNIQUE (handle),
	CONSTRAINT users_id PRIMARY KEY (uid)
);