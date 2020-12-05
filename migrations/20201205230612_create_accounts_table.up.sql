CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE accounts (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    balance int DEFAULT 0 NOT NULL, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);