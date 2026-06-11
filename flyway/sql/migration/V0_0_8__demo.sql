CREATE SCHEMA IF NOT EXISTS marketplacedemo;

CREATE TABLE IF NOT EXISTS marketplacedemo.category (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS marketplacedemo.products (
    id INTEGER PRIMARY KEY,
    title TEXT,
    description TEXT,
    price NUMERIC,
    discount_percentage NUMERIC,
    rating NUMERIC,
    stock INTEGER,
    brand TEXT,
    categoryId INTEGER REFERENCES marketplacedemo.category(id),
    thumbnail TEXT,
    search_vector tsvector,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS marketplacedemo.product_images (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES marketplacedemo.products(id),
    url TEXT
);

CREATE INDEX idx_product_search_fts_gin 
ON marketplacedemo.products USING GIN (search_vector) 
WHERE deleted_at IS NULL;

CREATE TABLE marketplacedemo.products_stage (
    id INTEGER PRIMARY KEY,
    title TEXT,
    description TEXT,
    price NUMERIC,
    discount_percentage NUMERIC,
    rating NUMERIC,
    stock INTEGER,
    brand TEXT,
    category TEXT,
    thumbnail TEXT,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS marketplacedemo.product_images_stage (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES marketplacedemo.products_stage(id),
    url TEXT
);

CREATE OR REPLACE FUNCTION marketplacedemo.normalize_marketplace()
RETURNS void AS $$
BEGIN
    INSERT INTO marketplacedemo.category (name)
    SELECT DISTINCT ps.category
    FROM marketplacedemo.products_stage ps
    WHERE ps.category IS NOT NULL
    ON CONFLICT (name) DO NOTHING;

    INSERT INTO marketplacedemo.products (
        id,
        title,
        description,
        price,
        discount_percentage,
        rating,
        stock,
        brand,
        categoryId,
        thumbnail,
        search_vector,
        deleted_at
    )
    SELECT
        ps.id,
        ps.title,
        ps.description,
        ps.price,
        ps.discount_percentage,
        ps.rating,
        ps.stock,
        ps.brand,
        c.id,
        ps.thumbnail,

        setweight(to_tsvector('english', coalesce(ps.title, '')), 'A') ||
        setweight(to_tsvector('english', coalesce(ps.description, '')), 'B') ||
        setweight(to_tsvector('english', coalesce(ps.brand, '')), 'C') ||
        setweight(to_tsvector('english', coalesce(c.name, '')), 'C'),

        ps.deleted_at

    FROM marketplacedemo.products_stage ps
    JOIN marketplacedemo.category c
      ON c.name = ps.category

    ON CONFLICT (id) DO UPDATE SET
        title = EXCLUDED.title,
        description = EXCLUDED.description,
        price = EXCLUDED.price,
        discount_percentage = EXCLUDED.discount_percentage,
        rating = EXCLUDED.rating,
        stock = EXCLUDED.stock,
        brand = EXCLUDED.brand,
        categoryId = EXCLUDED.categoryId,
        thumbnail = EXCLUDED.thumbnail,
        search_vector = EXCLUDED.search_vector,
        deleted_at = EXCLUDED.deleted_at;

    DELETE FROM marketplacedemo.product_images pi
    USING marketplacedemo.products_stage ps
    WHERE pi.product_id = ps.id;

    INSERT INTO marketplacedemo.product_images (product_id, url)
    SELECT
        pis.product_id,
        pis.url
    FROM marketplacedemo.product_images_stage pis
    JOIN marketplacedemo.products p
      ON p.id = pis.product_id;

END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS marketplacedemo.posts (
    id              UUID PRIMARY KEY DEFAULT extension.uuid_generate_v4(),
    content         TEXT NOT NULL,
    content_tsv     tsvector GENERATED ALWAYS AS (
        to_tsvector('spanish', content)
    ) STORED,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    product_id      INTEGER NOT NULL,
    reply_to_id     UUID NULL,
    user_id         UUID NOT NULL,
    CONSTRAINT fk_post_product
        FOREIGN KEY (product_id)
        REFERENCES marketplacedemo.products(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_reply_post
        FOREIGN KEY (reply_to_id)
        REFERENCES marketplacedemo.posts(id)
        ON DELETE SET NULL
);

CREATE INDEX idx_posts_product_id
ON marketplacedemo.posts(product_id);

CREATE INDEX idx_posts_content_tsv
ON marketplacedemo.posts
USING GIN(content_tsv);