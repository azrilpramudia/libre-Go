CREATE TABLE books (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title           VARCHAR(255) NOT NULL,
    isbn            VARCHAR(20) UNIQUE,
    openlibrary_id  VARCHAR(50),
    cover_url       TEXT,
    category_id     UUID REFERENCES categories(id) ON DELETE SET NULL,
    stock           INT NOT NULL DEFAULT 1 CHECK (stock >= 0),
    available       INT NOT NULL DEFAULT 1 CHECK (available >= 0),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_books_title ON books (title);