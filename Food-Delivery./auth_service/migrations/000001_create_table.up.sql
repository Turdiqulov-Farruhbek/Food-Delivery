create type privacy as enum ('public', 'private', 'contacts');
create type notification as enum ('yes', 'no');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(230) NOT NULL,
    email VARCHAR(230) NOT NULL unique,
    password VARCHAR(230) NOT NULL,
    full_name VARCHAR(230) not NULL,
    dob DATE,
    role VARCHAR(90),
    privacy_level privacy,
    theme VARCHAR(90),
    notification_enable notification,
    language VARCHAR,
    confirmed bool,
    working VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT DEFAULT 0
);