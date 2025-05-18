CREATE TABLE package_pickups (
    id UUID PRIMARY KEY,
    package_id UUID NOT NULL REFERENCES packages(id),
    locker_id UUID NOT NULL REFERENCES lockers(id),
    pickup_code VARCHAR(50) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
); 