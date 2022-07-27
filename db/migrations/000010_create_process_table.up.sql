CREATE TABLE IF NOT EXISTS process
(
    id          uuid PRIMARY KEY,
    created_at  timestamp                   NOT NULL,
    updated_at  timestamp,
    deleted_at  timestamp,
    winery_id   uuid references winery (id) NOT NULL,
    start_time  timestamp                   NOT NULL,
    end_time    timestamp,
    hash        varchar(100) UNIQUE,
    p_type      varchar(20)                 NOT NULL,
    transaction varchar(100),
    previous_id uuid references process (id)
);