CREATE DATABASE reservation
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'C'
    LC_CTYPE = 'C'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS reservation
(
    id integer NOT NULL DEFAULT 'nextval('reservation_id_seq'::regclass)',
    customer_id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    reservation_number integer NOT NULL,
    creation_date date,
    CONSTRAINT reservation_pkey PRIMARY KEY (id)
);