CREATE TABLE IF NOT EXISTS bottle
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL
);