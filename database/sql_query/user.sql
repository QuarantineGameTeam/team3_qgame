CREATE TABLE users
(
    id   INTEGER NOT NULL,
    name VARCHAR(30),
    team TEXT,
    role TEXT,
    health DECIMAL NOT NULL DEFAULT 100,
    strength DECIMAL NOT NULL DEFAULT 1,
    defence DECIMAL NOT NULL DEFAULT 1,
    intellect DECIMAL NOT NULL DEFAULT 1,
    level DECIMAL NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

INSERT INTO public.users
(name)
VALUES
('Luke'),
('Leia'),
('Han');