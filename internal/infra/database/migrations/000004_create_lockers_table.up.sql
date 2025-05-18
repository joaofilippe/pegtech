CREATE TABLE lockers (
    id UUID PRIMARY KEY,
    number VARCHAR(50) NOT NULL UNIQUE,
    size VARCHAR(50) NOT NULL,
    location VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    client_id UUID REFERENCES clients(id),
    reserved_at TIMESTAMP,
    occupied_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
); 