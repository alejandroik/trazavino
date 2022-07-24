CREATE TABLE winery
(
    id         uuid PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE cask
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL,
    c_type     varchar(20)                 NOT NULL,
    is_empty   bool                        NOT NULL
);

CREATE TABLE truck
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL
);

CREATE TABLE warehouse
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL,
    is_empty   bool                        NOT NULL
);

CREATE TABLE tank
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL,
    is_empty   bool                        NOT NULL
);

CREATE TABLE grape_type
(
    id         uuid PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE vineyard
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL
);

CREATE TABLE vineyard_grape_type
(
    vineyard_id   uuid references vineyard (id),
    grape_type_id uuid references grape_type (id),
    PRIMARY KEY (grape_type_id, vineyard_id)
);

CREATE TABLE wine
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL
);

CREATE TABLE process
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

CREATE TABLE reception
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

CREATE TABLE maceration
(
    id           uuid PRIMARY KEY references process (id),
    created_at   timestamp                      NOT NULL,
    updated_at   timestamp,
    deleted_at   timestamp,
    reception_id uuid references reception (id) NOT NULL,
    warehouse_id uuid references warehouse (id) NOT NULL
);

CREATE TABLE fermentation
(
    id           uuid PRIMARY KEY references process (id),
    created_at   timestamp                      NOT NULL,
    updated_at   timestamp,
    deleted_at   timestamp,
    warehouse_id uuid references warehouse (id) NOT NULL,
    tank_id      uuid references tank (id)      NOT NULL
);

CREATE TABLE ageing
(
    id         uuid PRIMARY KEY references process (id),
    created_at timestamp                 NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    tank_id    uuid references tank (id) NOT NULL,
    cask_id    uuid references cask (id) NOT NULL
);

CREATE TABLE bottling
(
    id         uuid PRIMARY KEY references process (id),
    created_at timestamp                 NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    cask_id    uuid references cask (id) NOT NULL,
    bottle_qty integer                   NOT NULL,
    wine_id    uuid references wine (id) NOT NULL
);

CREATE TABLE bottle
(
    id          uuid PRIMARY KEY,
    created_at  timestamp                     NOT NULL,
    updated_at  timestamp,
    deleted_at  timestamp,
    bottling_id uuid references bottling (id) NOT NULL,
    name        varchar(20)                   NOT NULL
);