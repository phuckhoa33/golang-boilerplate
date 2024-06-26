CREATE TABLE user_devices (
    id SERIAL PRIMARY KEY,
    user_id INT,
    device_id VARCHAR(255),
    access_type VARCHAR(255),
    browser VARCHAR(255),
    os VARCHAR(255),
    device_status VARCHAR(255),
    refresh_token VARCHAR(255),
    FOREIGN KEY (userId) REFERENCES users(id)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);