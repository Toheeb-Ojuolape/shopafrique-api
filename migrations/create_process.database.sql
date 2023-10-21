CREATE TABLE processes (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    email VARCHAR(255),
    process VARCHAR(255),
    expiry TIMESTAMP
);
