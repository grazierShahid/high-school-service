CREATE TABLE IF NOT EXISTS teacher_salary (
    id BIGSERIAL PRIMARY KEY,
    teacher_id INT,
    monthly_salary FLOAT,
    basic FLOAT,
    home_rent FLOAT,
    medical FLOAT,
    starting FLOAT,
    increment INT
);

--CREATE INDEX idx_teacher_salary_id ON teacher_salary(id);