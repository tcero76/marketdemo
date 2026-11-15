CREATE SCHEMA IF NOT EXISTS inventory;

CREATE TABLE IF NOT EXISTS inventory.category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS inventory.products (
    id UUID PRIMARY KEY,
    sku VARCHAR(100) NOT NULL,
    title TEXT,
    description TEXT,
    price NUMERIC(12, 2) NOT NULL,
    discount_percentage NUMERIC,
    rating NUMERIC,
    stock INTEGER,
    brand TEXT,
    categoryId UUID REFERENCES inventory.category(id),
    thumbnail TEXT,
    search_vector tsvector,
    deleted_at TIMESTAMPTZ,
    active      BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT products_price_positive
        CHECK (price > 0),
    CONSTRAINT products_sku_unique
        UNIQUE (sku)
);

CREATE TABLE IF NOT EXISTS inventory.product_images (
    id SERIAL PRIMARY KEY,
    product_id UUID REFERENCES inventory.products(id),
    url TEXT
);

CREATE INDEX idx_product_search_fts_gin 
ON inventory.products USING GIN (search_vector) 
WHERE deleted_at IS NULL;

CREATE TABLE inventory.stock (
    product_id  UUID PRIMARY KEY REFERENCES inventory.products(id),

    quantity    INTEGER NOT NULL DEFAULT 0,

    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT stock_quantity_non_negative
        CHECK (quantity >= 0)
);


CREATE TABLE inventory.inventory_movements (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    product_id  UUID NOT NULL REFERENCES inventory.products(id),
    user_id     UUID NOT NULL REFERENCES marketplace.users(user_id),

    movement_type VARCHAR(20) NOT NULL,

    quantity    INTEGER NOT NULL,

    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT inventory_movements_type_valid
        CHECK (
            movement_type IN (
                'PURCHASE',
                'SALE',
                'ADJUSTMENT'
            )
        ),

    CONSTRAINT inventory_movements_quantity_positive
        CHECK (quantity > 0)
);

CREATE INDEX idx_inventory_movements_product
    ON inventory.inventory_movements (product_id, created_at);

CREATE INDEX idx_inventory_movements_user
    ON inventory.inventory_movements (user_id);