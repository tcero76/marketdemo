CREATE SCHEMA IF NOT EXISTS extension;

ALTER SCHEMA extension OWNER TO tcero;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA extension;

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;

ALTER ROLE tcero SET search_path TO extension, marketplace, scrap, hydra, marketplacedemo, chat, public;