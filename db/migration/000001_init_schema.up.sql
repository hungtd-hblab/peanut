CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL,
	username text NULL,
	display_name text NULL,
	email text NULL,
	password text NULL,
	roles text NULL,
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
