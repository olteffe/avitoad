-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Europe/Moscow";

-- Create ads table
CREATE TABLE IF NOT EXISTS ads(
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name varchar(200),
    about text,
    photos text,
    price int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW ()
);

-- Create new ads
INSERT INTO ads (id, name, about, photos, price, created_at)
VALUES ('80c95ab0-a32e-4f13-9255-e83243e5ddb7', 'name1', 'about text1', 'http://example.com/11', 1000, '2011-10-19 10:23:54+02'),
       ('dc76012d-a15a-4615-9947-88ee7c791586', 'name2', 'about text2', 'http://example.com/21', 2000, '2012-10-19 10:23:54+02'),
       ('389390f8-e45f-45ad-91a0-06dbc8295966', 'name3', 'about text3', 'http://example.com/31', 3000, '2013-10-19 10:23:54+02');

-- Add indexes
CREATE INDEX time_ads_pagination ON ads (created_at, id);
CREATE INDEX price_ads_pagination ON ads (created_at, id);
