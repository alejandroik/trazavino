CREATE TABLE process
(
    id          bigserial PRIMARY KEY,
    created_at  timestamp    NOT NULL,
    updated_at  timestamp,
    deleted_at  timestamp,
    start_date  timestamp    NOT NULL,
    end_date    timestamp,
    hash        varchar(100) UNIQUE,
    p_type      varchar(20)  NOT NULL,
    transaction varchar(100),
    previous_id bigint
);

CREATE TABLE reception
(
    id         bigserial PRIMARY KEY references process (id),
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    weight     integer   NOT NULL,
    sugar      integer   NOT NULL
);