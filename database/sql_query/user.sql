CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id   uuid NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(30),
    team VARCHAR(30),
    role TEXT,
    health DECIMAL NOT NULL DEFAULT 100,
    strength DECIMAL NOT NULL DEFAULT 1,
    protection DECIMAL NOT NULL DEFAULT 1,
    intellect DECIMAL NOT NULL DEFAULT 1,
    level DECIMAL NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);