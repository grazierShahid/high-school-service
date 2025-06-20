CREATE TABLE IF NOT EXISTS class (
    id BIGSERIAL PRIMARY KEY,
    class_name TEXT,
    section_id INT,
    subjects TEXT[]
);

--CREATE INDEX idx_class_id ON class(id);