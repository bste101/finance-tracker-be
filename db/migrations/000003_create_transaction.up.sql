CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    category_id BIGINT NOT NULL
        REFERENCES categories(id)
        ON DELETE RESTRICT,

    type TEXT NOT NULL
        CHECK (type IN ('income', 'expense')),

    amount NUMERIC(12,2) NOT NULL
        CHECK (amount > 0),

    note TEXT,

    transaction_date DATE NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);