CREATE TABLE IF NOT EXISTS fermentation
(
    id           uuid PRIMARY KEY references process (id),
    created_at   timestamp                      NOT NULL,
    updated_at   timestamp,
    deleted_at   timestamp,
    warehouse_id uuid references warehouse (id) NOT NULL,
    tank_id      uuid references tank (id)      NOT NULL
);