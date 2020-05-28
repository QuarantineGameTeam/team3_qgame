-- Tickets
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id   uuid DEFAULT NOT NULL,
    name VARCHAR(30),
    PRIMARY KEY (id)
);
