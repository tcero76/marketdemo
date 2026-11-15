CREATE SCHEMA IF NOT EXISTS posts;

CREATE TABLE IF NOT EXISTS posts.posts (
    id              UUID PRIMARY KEY DEFAULT extension.uuid_generate_v4(),
    content         TEXT NOT NULL,
    content_tsv     tsvector GENERATED ALWAYS AS (
        to_tsvector('spanish', content)
    ) STORED,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    product_id      UUID NOT NULL,
    reply_to_id     UUID NULL,
    user_id         UUID NOT NULL,
    CONSTRAINT fk_post_product
        FOREIGN KEY (product_id)
        REFERENCES inventory.products(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_reply_post
        FOREIGN KEY (reply_to_id)
        REFERENCES posts.posts(id)
        ON DELETE SET NULL
);

CREATE INDEX idx_posts_product_id ON posts.posts(product_id);

CREATE INDEX idx_posts_content_tsv ON posts.posts USING GIN(content_tsv);