CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE credits (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    account_id uuid NOT NULL,
    credit_amount int NOT NULL, 
    available_amount int DEFAULT 0, 
    exausted boolean DEFAULT false, 
    type VARCHAR(32), 
    priority int, 
    expiry int, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY credits
    ADD CONSTRAINT credits_pkey PRIMARY KEY (id);