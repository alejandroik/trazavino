CREATE TABLE IF NOT EXISTS bottling
(
    id         uuid PRIMARY KEY references process (id),
    created_at timestamp                 NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    cask_id    uuid references cask (id) NOT NULL,
    bottle_qty integer                   NOT NULL,
    wine_id    uuid references wine (id) NOT NULL
);