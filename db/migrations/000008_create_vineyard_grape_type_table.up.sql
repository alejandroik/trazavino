CREATE TABLE IF NOT EXISTS vineyard_grape_type
(
    vineyard_id   uuid references vineyard (id),
    grape_type_id uuid references grape_type (id),
    PRIMARY KEY (grape_type_id, vineyard_id)
);