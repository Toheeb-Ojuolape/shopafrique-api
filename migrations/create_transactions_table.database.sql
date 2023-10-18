CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE transactions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    customer_email VARCHAR(255),
    customer_name VARCHAR(255),
    user_id DEFAULT gen_random_uuid(),
    payment_method VARCHAR(255),
    status VARCHAR(255),
    amount NUMERIC(10,2),
    type VARCHAR(255),
);