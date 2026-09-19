CREATE TABLE members (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name   VARCHAR(150) NOT NULL,
    email       VARCHAR(150) NOT NULL UNIQUE,
    phone       VARCHAR(20),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);