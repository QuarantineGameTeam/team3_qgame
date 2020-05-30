CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id   uuid not null DEFAULT uuid_generate_v4(),
    name VARCHAR(30),
    PRIMARY KEY (id)
);
