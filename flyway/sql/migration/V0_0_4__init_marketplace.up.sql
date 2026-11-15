--
-- Name: marketplace; Type: SCHEMA; Schema: -; Owner: tcero
--

CREATE SCHEMA IF NOT EXISTS marketplace;

ALTER SCHEMA marketplace OWNER TO tcero;

--
-- Name: productos; Type: TABLE; Schema: marketplace; Owner: tcero
--

CREATE TYPE marketplace.ciudad_enum AS ENUM ('Santiago','Viña del Mar','Antofagasta','Calama','La Serena','Concepción','Valdivia','Puerto Montt','Rancagua','Iquique','Temuco','Osorno','Copiapó','Arica','Talca','Punta Arenas','Chillán','Valparaíso','Curicó','Los Ángeles','Quilpué','Concon');

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
