CREATE TABLE employees (
    id UUID PRIMARY KEY REFERENCES users(id),
    role VARCHAR(50) NOT NULL
); 