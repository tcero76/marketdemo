
CREATE SCHEMA IF NOT EXISTS purchasing;

-- =========================================
-- PURCHASES
-- =========================================

CREATE TABLE purchasing.purchases (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    supplier_id     UUID NOT NULL REFERENCES marketplace.users(user_id),
    user_id         UUID NOT NULL REFERENCES marketplace.users(user_id),

    status          VARCHAR(20) NOT NULL DEFAULT 'PENDING',

    total_amount    NUMERIC(12, 2) NOT NULL,
    currency        CHAR(3) NOT NULL,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT purchases_total_positive
        CHECK (total_amount >= 0),

    CONSTRAINT purchases_currency_valid
        CHECK (currency ~ '^[A-Z]{3}$'),

    CONSTRAINT purchases_status_valid
        CHECK (
            status IN (
                'PENDING',
                'CANCELLED',
                'RECEIVED'
            )
        )
);

CREATE INDEX idx_purchases_supplier_id
    ON purchasing.purchases (supplier_id);

CREATE INDEX idx_purchases_user_id
    ON purchasing.purchases (user_id);

CREATE INDEX idx_purchases_status
    ON purchasing.purchases (status);

CREATE INDEX idx_purchases_created_at
    ON purchasing.purchases (created_at);


-- =========================================
-- PURCHASE ITEMS
-- =========================================

CREATE TABLE purchasing.purchase_items (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    purchase_id     UUID NOT NULL
        REFERENCES purchasing.purchases(id),

    product_id      UUID NOT NULL
        REFERENCES inventory.products(id),

    quantity        INTEGER NOT NULL DEFAULT 1,

    unit_price      NUMERIC(12, 2) NOT NULL,
    total_amount    NUMERIC(12, 2) NOT NULL,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT purchase_items_quantity_positive
        CHECK (quantity > 0),

    CONSTRAINT purchase_items_unit_price_positive
        CHECK (unit_price > 0),

    CONSTRAINT purchase_items_total_positive
        CHECK (total_amount > 0)
);

CREATE INDEX idx_purchase_items_purchase_id
    ON purchasing.purchase_items (purchase_id);

CREATE INDEX idx_purchase_items_product_id
    ON purchasing.purchase_items (product_id);


-- =========================================
-- RETURN PURCHASES
-- =========================================

CREATE TABLE purchasing.return_purchases (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    purchase_id     UUID NOT NULL
        REFERENCES purchasing.purchases(id),

    user_id         UUID NOT NULL
        REFERENCES marketplace.users(user_id),

    status          VARCHAR(20) NOT NULL DEFAULT 'PENDING',

    total_amount    NUMERIC(12, 2) NOT NULL,
    currency        CHAR(3) NOT NULL,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT return_purchases_total_positive
        CHECK (total_amount > 0),

    CONSTRAINT return_purchases_currency_valid
        CHECK (currency ~ '^[A-Z]{3}$'),

    CONSTRAINT return_purchases_status_valid
        CHECK (
            status IN (
                'PENDING',
                'CANCELLED',
                'COMPLETED'
            )
        )
);

CREATE INDEX idx_return_purchases_purchase_id
    ON purchasing.return_purchases (purchase_id);

CREATE INDEX idx_return_purchases_user_id
    ON purchasing.return_purchases (user_id);

CREATE INDEX idx_return_purchases_status
    ON purchasing.return_purchases (status);

CREATE TABLE purchasing.return_purchase_items (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    return_purchase_id  UUID NOT NULL
        REFERENCES purchasing.return_purchases(id),

    purchase_item_id    UUID NOT NULL
        REFERENCES purchasing.purchase_items(id),

    quantity            INTEGER NOT NULL,

    unit_price          NUMERIC(12, 2) NOT NULL,
    total_amount        NUMERIC(12, 2) NOT NULL,

    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT return_purchase_items_quantity_positive
        CHECK (quantity > 0),

    CONSTRAINT return_purchase_items_unit_price_positive
        CHECK (unit_price > 0),

    CONSTRAINT return_purchase_items_total_positive
        CHECK (total_amount > 0)
);

CREATE INDEX idx_return_purchase_items_return_purchase_id
    ON purchasing.return_purchase_items (return_purchase_id);

CREATE INDEX idx_return_purchase_items_purchase_item_id
    ON purchasing.return_purchase_items (purchase_item_id);

ALTER TABLE inventory.inventory_movements
ADD COLUMN purchase_id UUID
    REFERENCES purchasing.purchases(id);

ALTER TABLE inventory.inventory_movements
ADD COLUMN return_purchase_id UUID
    REFERENCES purchasing.return_purchases(id);