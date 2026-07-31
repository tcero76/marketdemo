--
-- Name: marketplace; Type: SCHEMA; Schema: -; Owner: tcero
--

CREATE SCHEMA IF NOT EXISTS marketplace;

ALTER SCHEMA marketplace OWNER TO tcero;

--
-- Name: productos; Type: TABLE; Schema: marketplace; Owner: tcero
--

CREATE TYPE marketplace.ciudad_enum AS ENUM ('Santiago','Viña del Mar','Antofagasta','Calama','La Serena','Concepción','Valdivia','Puerto Montt','Rancagua','Iquique','Temuco','Osorno','Copiapó','Arica','Talca','Punta Arenas','Chillán','Valparaíso','Curicó','Los Ángeles','Quilpué','Concon');

CREATE TABLE marketplace.productos (
    id integer,
    text text,
    author character varying(125),
    tags text,
    job_id uuid,
    processed_at timestamp without time zone
);

CREATE TYPE marketplace.OAuthProvider AS ENUM ('google','internal');

CREATE TABLE marketplace.users (
    user_id uuid,
    nombre text,
    password text,
    roles text,
    avatar text,
    tokenProvider text,
    provider marketplace.OAuthProvider,
    email VARCHAR(255),
    ID_provider VARCHAR(255) NULL,
    confirmado BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS marketplace.modelos (
	id integer,
    id_job bigint,
	modelo varchar(125) UNIQUE,
	descripcion text,
	created_at timestamp,
	deleted_at timestamptz null,
	PRIMARY KEY(id,id_job)
);

CREATE TABLE IF NOT EXISTS marketplace.posts (
	id integer,
	id_modelos integer,
    id_job bigint,
	descripcion text,
	modelo varchar(125),
	fechaRegistro timestamp,
	created_at timestamp,
	likes integer,
	PRIMARY KEY (id,id_job)
);

CREATE TABLE marketplace.outbox (
    id BIGSERIAL PRIMARY KEY,
    aggregate_type TEXT NOT NULL,
    aggregate_id TEXT NOT NULL,
    event_type TEXT NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_outbox_processed_created_at
ON marketplace.outbox (processed, created_at);

CREATE TABLE marketplace.posteos (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid,
  texto TEXT,
  menciones TEXT[],
  eliminado_en TIMESTAMPTZ,
  creado_en TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_posteos_menciones ON marketplace.posteos USING GIN (menciones);

CREATE TABLE marketplace.jwk_keys (
    kid        TEXT PRIMARY KEY,
    kty        TEXT NOT NULL,
    alg        TEXT NOT NULL,
    use        TEXT NOT NULL,
    private_jwk JSONB NOT NULL,
    public_jwk  JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP
);
