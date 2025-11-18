CREATE TABLE IF NOT EXISTS numbers (
    id SERIAL PRIMARY KEY,
    value INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_numbers_value ON numbers(value);
