CREATE TABLE IF NOT EXISTS reception
(
    id            uuid PRIMARY KEY references process (id),
    created_at    timestamp                       NOT NULL,
    updated_at    timestamp,
    deleted_at    timestamp,
    weight        integer                         NOT NULL,
    sugar         integer                         NOT NULL,
    truck_id      uuid references truck (id)      NOT NULL,
    vineyard_id   uuid references vineyard (id)   NOT NULL,
    grape_type_id uuid references grape_type (id) NOT NULL
);