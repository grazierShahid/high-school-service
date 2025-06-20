CREATE TABLE IF NOT EXISTS cash_media(
    id BIGSERIAL PRIMARY KEY,
    media_name TEXT NOT NULL, -- bkash, rocket, nagad
    payment_added FLOAT,
    salary_substruct FLOAT,
    current_ammount FLOAT
);

--CREATE INDEX idx_cash_media_id ON cash_media(id);