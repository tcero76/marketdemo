--
-- Name: scrap; Type: SCHEMA; Schema: -; Owner: tcero
--

CREATE SCHEMA IF NOT EXISTS scrap;

ALTER SCHEMA scrap OWNER TO tcero;

CREATE TABLE IF NOT EXISTS scrap.explore (
	id integer,
	id_modelos integer,
    id_job bigint not null,
	modelo varchar(125),
	fechaRegistro timestamp,
	created_at timestamp,
	likes integer,
	PRIMARY KEY (id,id_job)
);

CREATE TABLE IF NOT EXISTS scrap.modelos(
    id serial,
    id_job bigint not null,
    modelo character varying(125),
    descripcion text,
    fecharegistro timestamp,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS scrap.posts (
    id serial,
    id_modelos integer,
    id_job bigint not null,
    descripcion text,
    modelo character varying(125),
    fecharegistro timestamp,
    created_at timestamp,
    likes integer,
    primary key (id)
);

CREATE TABLE scrap.ts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_job bigint NOT NULL,
    portal TEXT NOT NULL,
    idpagina TEXT NOT NULL,
    nombre TEXT,
    edad INT,
    ciudad TEXT,
    servicios text[],
    servicios_adicionales text[],
    scraped_at TIMESTAMP DEFAULT NOW(),
    descripcion TEXT,
    UNIQUE (portal, idpagina, id_job)
);

CREATE TABLE IF NOT EXISTS scrap.servicios_alias (
    id SERIAL PRIMARY KEY,
    alias TEXT NOT NULL,
    id_servicio INT NOT NULL
);

CREATE TABLE scrap.modelos_instagram (
    id SERIAL PRIMARY KEY,
    modelo VARCHAR(255) NOT NULL,
    fecha_creacion TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    id_job BIGINT
);
