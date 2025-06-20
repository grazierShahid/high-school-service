CREATE TABLE IF NOT EXISTS classes (
    id BIGSERIAL PRIMARY KEY,
    class_id INT,
    class_teacher_id INT
);

--CREATE INDEX idx_classes_id ON classes(id);