CREATE DATABASE postdb;

USE postdb;

CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created DATETIME,
    updated DATETIME
);