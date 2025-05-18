CREATE TABLE packages (
    id UUID PRIMARY KEY,
    tracking_code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    recipient_id UUID NOT NULL REFERENCES clients(id),
    locker_id UUID REFERENCES lockers(id),
    pickup_password VARCHAR(50) NOT NULL,
    pickup_expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
); 