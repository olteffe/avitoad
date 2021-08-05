BEGIN;
-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Europe/Moscow";

-- Del table
DROP TABLE IF EXISTS "ads";

-- Drop Index
DROP INDEX IF EXISTS time_ads_pagination, price_ads_pagination;

-- Create ads table
CREATE TABLE IF NOT EXISTS ads(
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name varchar(200),
    about varchar(1000),
    photos text[3],
    price int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    first_photo text
);

-- Create new ads
INSERT INTO ads (id, name, about, photos, price, created_at, first_photo)
VALUES ('80c95ab0-a32e-4f13-9255-e83243e5ddb7', 'name1', 'about text1', '{"http://example.com/11", "http://example.com/12", "http://example.com/13"}', 1000, '2011-10-19 10:23:54+02', 'http://example.com/11'),
       ('dc76012d-a15a-4615-9947-88ee7c791586', 'name2', 'about text2', '{"http://example.com/21", "http://example.com/22", "http://example.com/23"}', 2000, '2012-10-19 10:23:54+02', 'http://example.com/21'),
       ('389390f8-e45f-45ad-91a0-06dbc8295966', 'name3', 'about text3', '{"http://example.com/31", "http://example.com/32", "http://example.com/33"}', 3000, '2013-10-19 10:23:54+02', 'http://example.com/21');

-- Add indexes
CREATE INDEX time_ads_pagination ON ads (created_at, id);
CREATE INDEX price_ads_pagination ON ads (created_at, id);

END;
