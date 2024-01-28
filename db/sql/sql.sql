-- Postgresql

-- @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>

-- Create Database
SELECT 'CREATE DATABASE mypriv_go'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mypriv_go')\gexec

-- Connect to Database
\connect mypriv_go

-- Sequences
CREATE SEQUENCE IF NOT EXISTS ip_address_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS user_accounts_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS personal_information_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS follows_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS subscriptions_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS routes_defautls_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS dashboard_data_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

-- Ip Address Access
CREATE TABLE IF NOT EXISTS ip_address_access (
    id BIGINT PRIMARY KEY DEFAULT nextval('ip_address_id_seq'),
    ip_address VARCHAR(50) NOT NULL UNIQUE,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- User Accounts
CREATE TABLE IF NOT EXISTS user_accounts (
    id BIGINT PRIMARY KEY DEFAULT nextval('user_accounts_id_seq'),
    email TEXT NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    nickname VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Personal Information
CREATE TABLE IF NOT EXISTS personal_information (
    id BIGINT PRIMARY KEY DEFAULT nextval('personal_information_id_seq'),
    user_account_id BIGINT NOT NULL UNIQUE,
    first_name VARCHAR(50) NULL,
    last_name VARCHAR(50) NULL,
    profile_picture TEXT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id)
);

-- Follows USER_ACCOUNTS
CREATE TABLE IF NOT EXISTS follows (
    id BIGINT PRIMARY KEY DEFAULT nextval('follows_id_seq'),
    user_account_id BIGINT NOT NULL,
    follower_id BIGINT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES user_accounts(id) ON DELETE CASCADE,
    CONSTRAINT unique_follow_pair UNIQUE (user_account_id, follower_id)
);

-- Subscriptions per User followers 
CREATE TABLE IF NOT EXISTS subscriptions_followers (
    id BIGINT PRIMARY KEY DEFAULT nextval('subscriptions_id_seq'),
    user_account_id BIGINT NOT NULL,
    follower_id BIGINT NOT NULL,
    data JSONB NOT NULL,
    active_sub BOOLEAN NOT NULL DEFAULT TRUE,
    start_sub TIMESTAMP NOT NULL DEFAULT NOW(),
    finish_sub TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES user_accounts(id) ON DELETE CASCADE,
    CONSTRAINT unique_subscription_pair UNIQUE (user_account_id, follower_id)
);

-- Dashboard Posts
-- CREATE TABLE IF NOT EXISTS dashboard_posts (
--     id BIGINT PRIMARY KEY DEFAULT nextval('dashboard_posts_id_seq'),
--     user_account_id BIGINT NOT NULL,
--     post JSONB NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE
-- );

-- Dashboard Data per User
CREATE TABLE IF NOT EXISTS dashboard_data (
    id BIGINT PRIMARY KEY DEFAULT nextval('dashboard_data_id_seq'),
    user_account_id BIGINT NOT NULL,
    data JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE
);

CREATE TABLE  IF NOT EXISTS routes_defautls (
    id BIGINT PRIMARY KEY DEFAULT nextval('routes_defautls_id_seq'),
    route VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create Indexes
CREATE INDEX IF NOT EXISTS ip_address_access_ip_address_idx ON ip_address_access (ip_address);
CREATE INDEX IF NOT EXISTS user_accounts_email_idx ON user_accounts (email);
CREATE INDEX IF NOT EXISTS user_accounts_nickname_idx ON user_accounts (nickname);
CREATE INDEX IF NOT EXISTS personal_information_user_account_id_idx ON personal_information (user_account_id);
CREATE INDEX IF NOT EXISTS follows_user_account_id_idx ON follows (user_account_id);
CREATE INDEX IF NOT EXISTS follows_follower_id_idx ON follows (follower_id);
CREATE INDEX IF NOT EXISTS subscriptions_followers_user_account_id_idx ON subscriptions_followers (user_account_id);
CREATE INDEX IF NOT EXISTS subscriptions_followers_follower_id_idx ON subscriptions_followers (follower_id);
CREATE INDEX IF NOT EXISTS dashboard_data_user_account_id_idx ON dashboard_data (user_account_id);

-- PROCEDURES

-- Insert New Ip Address
CREATE OR REPLACE PROCEDURE insert_ip_address(ip_address_p TEXT)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO ip_address_access (ip_address, active) 
    VALUES (ip_address_p, TRUE); 
END;
$$;

-- Insert New User Account
CREATE OR REPLACE PROCEDURE insert_new_user_account(
    email_ TEXT, 
    password_ TEXT, 
    nickname_ TEXT,
    first_name_ TEXT,
    last_name_ TEXT,
    profile_picture_ TEXT
)
LANGUAGE plpgsql
AS $$
DECLARE
    new_user_id BIGINT;
BEGIN
    INSERT INTO user_accounts (email, password, nickname) 
    VALUES (email_, password_, nickname_)
    RETURNING id INTO new_user_id;

    INSERT INTO personal_information (user_account_id ,first_name, last_name, profile_picture)
    VALUES (new_user_id, first_name_, last_name_, profile_picture_);

    INSERT INTO dashboard_data (user_account_id, data) 
    VALUES (new_user_id, '{}'::jsonb);
END;
$$;

-- Update User Account
CREATE OR REPLACE PROCEDURE update_user_account(
    user_account_id BIGINT,
    email_param TEXT, 
    password_param TEXT, 
    nickname_param TEXT,
    first_name_param TEXT,
    last_name_param TEXT,
    profile_picture_param TEXT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE user_accounts 
    SET 
        email = COALESCE(email_param, email),
        password = COALESCE(password_param, password),
        nickname = COALESCE(nickname_param, nickname)
    WHERE id = user_account_id;

    UPDATE personal_information 
    SET 
        first_name = COALESCE(first_name_param, first_name),
        last_name = COALESCE(last_name_param, last_name),
        profile_picture = COALESCE(profile_picture_param, profile_picture)
    WHERE user_account_id = user_account_id;
END;
$$;

-- Insert New Follow
CREATE OR REPLACE PROCEDURE insert_row_follow(
    user_account_id_p BIGINT,
    follower_id_p BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO follows (user_account_id, follower_id) 
    VALUES (user_account_id_p, follower_id_p);
END;
$$;

-- Update Follow
CREATE OR REPLACE PROCEDURE update_row_follow(
    user_account_id_p BIGINT,
    follower_id_p BIGINT,
    active_p BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE follows 
    SET active = COALESCE(active_p)
    WHERE user_account_id = user_account_id_p 
    AND follower_id = follower_id_p;
END;
$$;

-- Insert New Subscription
CREATE OR REPLACE PROCEDURE insert_row_subscription(
    user_account_id_p BIGINT,
    follower_id_p BIGINT,
    data_p JSONB,
    active_sub_p BOOLEAN,
    start_sub_p TIMESTAMP,
    finish_sub_p TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO subscriptions_followers (user_account_id, follower_id, data, active_sub, start_sub, finish_sub) 
    VALUES (user_account_id_p, follower_id_p, data_p, active_sub_p, start_sub_p, finish_sub_p);
END;
$$;

-- Update Subscription
CREATE OR REPLACE PROCEDURE update_row_subscription(
    user_account_id_p BIGINT,
    follower_id_p BIGINT,
    data_p JSONB,
    active_sub_p BOOLEAN,
    start_sub_p TIMESTAMP,
    finish_sub_p TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE subscriptions_followers 
    SET 
        data = COALESCE(data_p, data),
        active_sub = COALESCE(active_sub_p, active_sub),
        start_sub = COALESCE(start_sub_p, start_sub),
        finish_sub = COALESCE(finish_sub_p, finish_sub)
    WHERE user_account_id = user_account_id_p 
    AND follower_id = follower_id_p;
END;
$$;

-- -- Insert New Dashboard Data
-- CREATE OR REPLACE PROCEDURE insert_row_dashboard_data(
--     user_account_id_p BIGINT,
--     data_p JSONB
-- )
-- LANGUAGE plpgsql
-- AS $$
-- BEGIN
--     INSERT INTO dashboard_data (user_account_id, data) 
--     VALUES (user_account_id_p, data_p);
-- END;


-- Update Dashboard Data
CREATE OR REPLACE PROCEDURE update_row_dashboard_data(
    user_account_id_p BIGINT,
    data_p JSONB
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE dashboard_data 
    SET data = COALESCE(data_p, data)
    WHERE user_account_id = user_account_id_p;
END;
$$;

-- FUNCTIONS - QUERYS

-- Get all data of user
CREATE OR REPLACE FUNCTION get_user_all_data(
    email_p TEXT
)
RETURNS TABLE (
    id_user BIGINT,
    nickname VARCHAR(50),
    created_at TIMESTAMP,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    profile_picture TEXT,
    data JSONB
)
AS $$
BEGIN
    RETURN QUERY
    SELECT t1.id AS id_user, t1.nickname, t1.created_at, t2.first_name, 
    t2.last_name, t2.profile_picture, t3.data
    FROM user_accounts AS t1
    INNER JOIN personal_information AS t2 ON t1.id = t2.user_account_id
    INNER JOIN dashboard_data AS t3 ON t1.id = t3.user_account_id
    WHERE t1.email = email_p;
END;
$$
LANGUAGE plpgsql;

-- Get Followers by user
CREATE OR REPLACE FUNCTION get_followers_by_user(
    user_account_id_p BIGINT
)
RETURNS TABLE (
    user_account_id BIGINT,
    follower_id BIGINT
)
AS $$
BEGIN
    RETURN QUERY
    SELECT 
    user_account_id, follower_id,
    FROM follows 
    WHERE user_account_id = user_account_id_p
    OR follower_id = user_account_id_p;
    AND active = TRUE;
END;
$$
LANGUAGE plpgsql;

-- Get Subscriptions by user
CREATE OR REPLACE FUNCTION get_subscriptions_by_user(
    user_account_id_p BIGINT
)
RETURNS TABLE (
    user_account_id BIGINT,
    follower_id BIGINT,
    data JSONB,
    start_sub TIMESTAMP,
    finish_sub TIMESTAMP
)
AS $$
BEGIN
    RETURN QUERY
    SELECT 
    user_account_id,follower_id,data,
    start_sub,finish_sub
    FROM subscriptions_followers 
    WHERE user_account_id = user_account_id_p
    AND active_sub = TRUE;
END;
$$
LANGUAGE plpgsql;

-- Get Dashboard Data by user
CREATE OR REPLACE FUNCTION get_dashboard_data_by_user(
    user_account_id_p BIGINT
)
RETURNS TABLE (
    data JSONB
)
AS $$
BEGIN
    RETURN QUERY
    SELECT data
    FROM dashboard_data 
    WHERE user_account_id = user_account_id_p;
END;
$$
LANGUAGE plpgsql;

-- Default Routes
CREATE OR REPLACE FUNCTION get_default_routes()
RETURNS TABLE (
    route_ VARCHAR(50)
)
AS $$
BEGIN
    RETURN QUERY
    SELECT route as route_
    FROM routes_defautls;
END;
$$
LANGUAGE plpgsql;

-- Validate if user exists
CREATE OR REPLACE FUNCTION validate_user_exists(
    email_p TEXT
)
RETURNS TABLE (
    email TEXT
)
AS $$
BEGIN
    RETURN QUERY
    SELECT email
    FROM user_accounts 
    WHERE email = email_p;
END;
$$
LANGUAGE plpgsql;

-- Insert Rows Values

INSERT INTO routes_defautls (route) 
values ('my-priv');

INSERT INTO ip_address_access (ip_address, active)
VALUES ('127.0.0.1', TRUE);