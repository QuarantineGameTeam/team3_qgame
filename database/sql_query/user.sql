CREATE TABLE users
(
    id        INTEGER   NOT NULL,
    name      VARCHAR(30),
    team      TEXT,
    status    TEXT      DEFAULT 'main',
    health    DECIMAL   NOT NULL DEFAULT 100,
    strength  DECIMAL   NOT NULL DEFAULT 1,
    defence   DECIMAL   NOT NULL DEFAULT 1,
    intellect DECIMAL   NOT NULL DEFAULT 1,
    level     DECIMAL   NOT NULL DEFAULT 0,
    currency  DECIMAL   NOT NULL DEFAULT 50,
    inventory INTEGER[] NOT NULL DEFAULT ARRAY []::integer[],
    PRIMARY KEY (id)
);
