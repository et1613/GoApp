-- Initial schema: core tables only (users, contacts, status, blocked_users)
-- Chat/messaging tables moved to migration 0003

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    display_name VARCHAR(100),
    profile_picture_url TEXT,
    about_text VARCHAR(250),
    last_seen_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE contacts (
    user_id uuid NOT NULL REFERENCES users(id),
    contact_phone_number VARCHAR(20) NOT NULL,
    contact_user_id uuid REFERENCES users(id),
    display_name_override VARCHAR(100),
    PRIMARY KEY (user_id, contact_phone_number)
);

CREATE TABLE status_updates (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES users(id),
    content_type VARCHAR(20) NOT NULL, -- 'text', 'image', 'video'
    content_or_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX ON status_updates (user_id, created_at);

CREATE TABLE status_views (
    status_id uuid NOT NULL REFERENCES status_updates(id),
    user_id uuid NOT NULL REFERENCES users(id),
    viewed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (status_id, user_id)
);

CREATE TABLE blocked_users (
    blocker_user_id uuid NOT NULL REFERENCES users(id),
    blocked_user_id uuid NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (blocker_user_id, blocked_user_id)
);

CREATE TABLE user_devices (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES users(id),
    refresh_token_hash VARCHAR(255) NOT NULL,
    device_name VARCHAR(100),
    device_type VARCHAR(20), -- 'mobile', 'web', 'desktop'
    push_notification_token TEXT,
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
