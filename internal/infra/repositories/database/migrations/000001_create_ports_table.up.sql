CREATE TABLE IF NOT EXISTS ports (
    id SERIAL PRIMARY KEY,
    locker INTEGER NOT NULL,
    port INTEGER NOT NULL,
    number VARCHAR(50) NOT NULL,
    package_code VARCHAR(100),
    package_pickup_password VARCHAR(100),
    package_pickup_expires_at TIMESTAMP,
    package_user_id UUID,
    status VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE',
    reserved_expiration TIMESTAMP,
    occupied_at TIMESTAMP,
    occupied_until TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_locker FOREIGN KEY (locker) REFERENCES lockers(id) ON DELETE CASCADE,
    CONSTRAINT unique_port_number UNIQUE (locker, port)
);

-- Create index for faster lookups
CREATE INDEX idx_ports_locker ON ports(locker);
CREATE INDEX idx_ports_status ON ports(status);
CREATE INDEX idx_ports_package_user ON ports(package_user_id);

-- Add trigger to automatically update updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_ports_updated_at
    BEFORE UPDATE ON ports
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column(); 