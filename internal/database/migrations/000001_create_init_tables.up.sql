-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE IF NOT EXISTS ads (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name varchar(200),
    about text,
    photos text[3],
    price numeric(12,2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW ()
);

-- Add indexes
--TODO later
