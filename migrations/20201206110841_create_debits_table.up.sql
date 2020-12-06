CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE debits (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    account_id uuid NOT NULL,
    amount int DEFAULT 0,
    used_credits int NOT NULL, 
    used_credit_ids uuid[] DEFAULT array[]::uuid[], 
    type VARCHAR(32), 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY debits
    ADD CONSTRAINT debits_pkey PRIMARY KEY (id);