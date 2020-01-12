DROP TABLE IF EXISTS ads;

CREATE TABLE IF NOT EXISTS ads
(
    id          bigserial                    not null primary key,
    name        varchar(200) COLLATE "POSIX" not null,
    created     timestamptz DEFAULT now(),
    price       float                        not null,
    description varchar(1000),
    photos      varchar[3]
);