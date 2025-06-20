CREATE TABLE IF NOT EXISTS teacher_background (
    id BIGSERIAL PRIMARY KEY,
    teacher_id INT,
    ex_schools TEXT[],
    education TEXT[]
);

--CREATE INDEX idx_teacher_id ON teacher_background(id);