CREATE TABLE IF NOT EXISTS maceration
(
    id           uuid PRIMARY KEY references process (id),
    created_at   timestamp                      NOT NULL,
    updated_at   timestamp,
    deleted_at   timestamp,
    reception_id uuid references reception (id) NOT NULL,
    warehouse_id uuid references warehouse (id) NOT NULL
);