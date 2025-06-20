CREATE TABLE IF NOT EXISTS invoice_bill(
    id BIGSERIAL PRIMARY KEY,
    student_id INT,
    student_bill_history_id INT,
    cash_media_id INT
);

--CREATE INDEX idx_invoice_bill_id ON invoice_bill(id);