CREATE TABLE IF NOT EXISTS student_background (
    id BIGSERIAL PRIMARY KEY,
    student_id INT,
    ex_schools TEXT[],
    results_history_id INT
);

--CREATE INDEX idx_student_background_id ON student_background(id);