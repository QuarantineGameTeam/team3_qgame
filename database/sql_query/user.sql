-- Tickets
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id   uuid NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(30),
    team TEXT,
    role TEXT,
    health DECIMAL NOT NULL DEFAULT 100,
    str DECIMAL NOT NULL DEFAULT 1,
    def DECIMAL NOT NULL DEFAULT 1,
    int DECIMAL NOT NULL DEFAULT 1,
    lvl DECIMAL NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);