CREATE TABLE loans (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id      UUID NOT NULL REFERENCES books(id) ON DELETE RESTRICT,
    member_id    UUID NOT NULL REFERENCES members(id) ON DELETE RESTRICT,
    borrowed_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    due_at       TIMESTAMPTZ NOT NULL,
    returned_at  TIMESTAMPTZ,
    status       VARCHAR(20) NOT NULL DEFAULT 'borrowed' CHECK (status IN ('borrowed', 'returned', 'overdue')),
    fine_amount  NUMERIC(10,2) NOT NULL DEFAULT 0
);

CREATE INDEX idx_loans_member ON loans (member_id);
CREATE INDEX idx_loans_book ON loans (book_id);
CREATE INDEX idx_loans_status ON loans (status);