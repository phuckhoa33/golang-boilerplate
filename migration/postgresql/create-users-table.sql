CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phoneNumber VARCHAR(20) NOT NULL,
    fullname VARCHAR(100) NOT NULL,
    address VARCHAR(200) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    dateOfBirth DATE NOT NULL,
    password VARCHAR(100) NOT NULL,
    avatar VARCHAR(200) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    deleted_at TIMESTAMP
    -- Create roles column here and add a foreign key constraint to the roles table
    roles TEXT[] NOT NULL,
    FOREIGN KEY (roles) REFERENCES roles(permissions)
);

-- Create two users for data sample
INSERT INTO users (username, email, phoneNumber, fullname, address, gender, dateOfBirth, password, avatar, roles)
VALUES ('admin', 'admin@gmail.com', '0123456789', 'Admin', 'Hanoi', 'MALE', '1990-01-01', 'admin', 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50', '{ADMIN}');
INSERT INTO users (username, email, phoneNumber, fullname, address, gender, dateOfBirth, password, avatar, roles)
VALUES ('phuckhoa', 'phuckhoa81@gmail.com', '0123456789', 'Phuc Khoa', 'Hanoi', 'MALE', '2003-03-03', 'phuckhoa', 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50', '{USER}');