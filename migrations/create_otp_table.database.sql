CREATE TABLE otps (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    email VARCHAR(255),
    otp VARCHAR(255),
    expired_at TIMESTAMP
);
