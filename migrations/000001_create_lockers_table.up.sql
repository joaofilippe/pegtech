CREATE TABLE lockers (
    id SERIAL PRIMARY KEY,
    number VARCHAR(10) NOT NULL,
    location VARCHAR(100) NOT NULL,
    package_code VARCHAR(50),
    package_pickup_password VARCHAR(6),
    package_pickup_expires_at TIMESTAMP,
    package_user_id UUID,
    status VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE',
    client_id UUID,
    reserved_expiration TIMESTAMP,
    occupied_at TIMESTAMP,
    occupied_until TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES users(id) ON DELETE SET NULL
);

-- Create index for faster lookups
CREATE INDEX idx_lockers_status ON lockers(status);
CREATE INDEX idx_lockers_package_user_id ON lockers(package_user_id);
CREATE INDEX idx_lockers_client_id ON lockers(client_id); 