CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    email VARCHAR(255) UNIQUE,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    country VARCHAR(255),
    phone_number VARCHAR(255) UNIQUE,
    business_name VARCHAR(255),
    business_type VARCHAR(255),
    lightning_address VARCHAR(255),
    password VARCHAR(255),
    balance VARCHAR(255),
    role VARCHAR(255)
);
