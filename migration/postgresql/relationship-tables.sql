
-- Drop the tables if they exist, in reverse order of their dependencies
DROP TABLE IF EXISTS user_devices CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


-- Recreate the roles table
CREATE TABLE roles (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    permissions TEXT[] NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO roles (id, name, permissions)
VALUES ('eca2a00d-485d-4c10-b4bd-a33f38fb53b7','ADMIN', '{"CREATE_USER", "READ_ALL_DATA", "UPDATE_ROLES"}');


CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    address VARCHAR(200) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    date_of_birth DATE NOT NULL,
    password VARCHAR(100) NOT NULL,
    avatar VARCHAR(200) NOT NULL,
    verify_account_otp VARCHAR(6),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    -- Create roles column here and add a foreign key constraint to the roles table
    role_id uuid NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE user_devices (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL,
    access_type VARCHAR(255),
    browser VARCHAR(255),
    os VARCHAR(255),
    device_status VARCHAR(255),
    refresh_token VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);