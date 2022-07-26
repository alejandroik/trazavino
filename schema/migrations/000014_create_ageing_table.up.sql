CREATE TABLE IF NOT EXISTS ageing
(
    id         uuid PRIMARY KEY references process (id),
    created_at timestamp                 NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    tank_id    uuid references tank (id) NOT NULL,
    cask_id    uuid references cask (id) NOT NULL
);