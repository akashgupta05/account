CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE transactions (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    account_id uuid NOT NULL,
    user_id uuid NOT NULL,
    amount int NOT NULL, 
    type VARCHAR(32), 
    priority int, 
    expiry int, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);