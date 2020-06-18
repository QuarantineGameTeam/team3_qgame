CREATE TABLE users
(
    id        INTEGER NOT NULL,
    name      VARCHAR(30),
    team      TEXT,
    status    bool    NOT NULL DEFAULT FALSE,
    health    DECIMAL NOT NULL DEFAULT 100,
    strength  DECIMAL NOT NULL DEFAULT 1,
    defence   DECIMAL NOT NULL DEFAULT 1,
    intellect DECIMAL NOT NULL DEFAULT 1,
    level     DECIMAL NOT NULL DEFAULT 0,
    inventory varchar(20)[],
    PRIMARY KEY (id)
);


