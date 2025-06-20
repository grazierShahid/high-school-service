CREATE TABLE IF NOT EXISTS results (
    id BIGSERIAL PRIMARY KEY,
    student_id INT,
    class_id INT,
    result_sheet_id INT
);

--CREATE INDEX idx_results_id ON results(id);