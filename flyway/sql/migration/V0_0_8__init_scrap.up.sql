--
-- Name: scrap; Type: SCHEMA; Schema: -; Owner: tcero
--

CREATE SCHEMA IF NOT EXISTS scrap;

ALTER SCHEMA scrap OWNER TO tcero;

CREATE TABLE scrap.products_stage (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sku VARCHAR(100) NOT NULL,
    title TEXT,
    description TEXT,
    price NUMERIC,
    discount_percentage NUMERIC,
    rating NUMERIC,
    quantity INTEGER,
    brand TEXT,
    category TEXT,
    thumbnail TEXT,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS scrap.product_images_stage (
    id SERIAL PRIMARY KEY,
    product_id UUID REFERENCES scrap.products_stage(id),
    url TEXT
);

CREATE OR REPLACE FUNCTION scrap.normalize_marketplace()
RETURNS void AS $$
BEGIN
    INSERT INTO inventory.category (name)
    SELECT DISTINCT ps.category
    FROM scrap.products_stage ps
    WHERE ps.category IS NOT NULL
    ON CONFLICT (name) DO NOTHING;

    INSERT INTO inventory.products (
        id,
        sku,
        title,
        description,
        price,
        discount_percentage,
        rating,
        quantity,
        brand,
        categoryId,
        thumbnail,
        search_vector,
        deleted_at
    )
    SELECT
        ps.id,
        ps.sku,
        ps.title,
        ps.description,
        ps.price,
        ps.discount_percentage,
        ps.rating,
        ps.quantity,
        ps.brand,
        c.id,
        ps.thumbnail,

        setweight(to_tsvector('english', coalesce(ps.title, '')), 'A') ||
        setweight(to_tsvector('english', coalesce(ps.description, '')), 'B') ||
        setweight(to_tsvector('english', coalesce(ps.brand, '')), 'C') ||
        setweight(to_tsvector('english', coalesce(c.name, '')), 'C'),

        ps.deleted_at

    FROM scrap.products_stage ps
    JOIN inventory.category c
      ON c.name = ps.category

    ON CONFLICT (id) DO UPDATE SET
        title = EXCLUDED.title,
        sku = EXCLUDED.sku,
        description = EXCLUDED.description,
        price = EXCLUDED.price,
        discount_percentage = EXCLUDED.discount_percentage,
        rating = EXCLUDED.rating,
        quantity = EXCLUDED.quantity,
        brand = EXCLUDED.brand,
        categoryId = EXCLUDED.categoryId,
        thumbnail = EXCLUDED.thumbnail,
        search_vector = EXCLUDED.search_vector,
        deleted_at = EXCLUDED.deleted_at;

    DELETE FROM inventory.product_images pi
    USING scrap.products_stage ps
    WHERE pi.product_id = ps.id;

    INSERT INTO inventory.product_images (product_id, url)
    SELECT
        pis.product_id,
        pis.url
    FROM scrap.product_images_stage pis
    JOIN inventory.products p
      ON p.id = pis.product_id;

END;
$$ LANGUAGE plpgsql;