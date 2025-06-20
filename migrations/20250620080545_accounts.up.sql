CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,
    balance FLOAT,
    plus FLOAT,
    minux FLOAT,
    transaction FLOAT
);

--CREATE INDEX idx_accounts_id ON accounts(id);