CREATE TABLE user_devices (
    id SERIAL PRIMARY KEY,
    userId INT,
    deviceId VARCHAR(255),
    accessType VARCHAR(255),
    browser VARCHAR(255),
    os VARCHAR(255),
    deviceStatus VARCHAR(255),
    refreshToken VARCHAR(255),
    FOREIGN KEY (userId) REFERENCES users(id)
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    deleted_at TIMESTAMP
);