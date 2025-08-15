CREATE TYPE transaction_state AS ENUM ('win', 'lose');
CREATE TYPE transaction_source_type AS ENUM ('game', 'server', 'payment');

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount DECIMAL(12, 2) NOT NULL,
    state transaction_state NOT NULL,
    source_type transaction_source_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
