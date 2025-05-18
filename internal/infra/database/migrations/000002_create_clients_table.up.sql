CREATE TABLE clients (
    id UUID PRIMARY KEY REFERENCES users(id),
    phone VARCHAR(50) NOT NULL,
    address TEXT NOT NULL
); 