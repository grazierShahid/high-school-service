CREATE TABLE IF NOT EXISTS payments(
    id BIGSERIAL PRIMARY KEY,
    student_id INT,
    invoice_id INT
);

--CREATE INDEX idx_payments_id ON paypemts(id);