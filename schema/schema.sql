CREATE TABLE cellar
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE bottle
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE cask
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL,
    c_type     varchar(20) NOT NULL,
    is_empty   bool        NOT NULL
);

CREATE TABLE cellar_cask
(
    cellar_id bigint references cellar (id),
    cask_id   bigint references cask (id),
    PRIMARY KEY (cellar_id, cask_id)
);

CREATE TABLE truck
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE cellar_truck
(
    cellar_id bigint references cellar (id),
    truck_id  bigint references truck (id),
    PRIMARY KEY (cellar_id, truck_id)
);

CREATE TABLE warehouse
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL,
    is_empty   bool        NOT NULL
);

CREATE TABLE cellar_warehouse
(
    cellar_id    bigint references cellar (id),
    warehouse_id bigint references warehouse (id),
    PRIMARY KEY (cellar_id, warehouse_id)
);

CREATE TABLE tank
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL,
    is_empty   bool        NOT NULL
);

CREATE TABLE cellar_tank
(
    cellar_id bigint references cellar (id),
    tank_id   bigint references tank (id),
    PRIMARY KEY (cellar_id, tank_id)
);

CREATE TABLE grape_type
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE vineyard
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE cellar_vineyard
(
    cellar_id   bigint references cellar (id),
    vineyard_id bigint references vineyard (id),
    PRIMARY KEY (cellar_id, vineyard_id)
);

CREATE TABLE vineyard_grape_type
(
    vineyard_id   bigint references vineyard (id),
    grape_type_id bigint references grape_type (id),
    PRIMARY KEY (grape_type_id, vineyard_id)
);

CREATE TABLE wine
(
    id         bigserial PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);

CREATE TABLE cellar_wine
(
    cellar_id bigint references cellar (id),
    wine_id   bigint references wine (id),
    PRIMARY KEY (cellar_id, wine_id)
);

CREATE TABLE process
(
    id          bigserial PRIMARY KEY,
    created_at  timestamp   NOT NULL,
    updated_at  timestamp,
    deleted_at  timestamp,
    start_date  timestamp   NOT NULL,
    end_date    timestamp,
    hash        varchar(100) UNIQUE,
    p_type      varchar(20) NOT NULL,
    transaction varchar(100),
    previous_id bigint
);

CREATE TABLE reception
(
    id            bigint PRIMARY KEY references process (id),
    created_at    timestamp                         NOT NULL,
    updated_at    timestamp,
    deleted_at    timestamp,
    weight        integer                           NOT NULL,
    sugar         integer                           NOT NULL,
    truck_id      bigint references truck (id)      NOT NULL,
    vineyard_id   bigint references vineyard (id)   NOT NULL,
    grape_type_id bigint references grape_type (id) NOT NULL
);

CREATE TABLE maceration
(
    id           bigint PRIMARY KEY references process (id),
    created_at   timestamp                        NOT NULL,
    updated_at   timestamp,
    deleted_at   timestamp,
    reception_id bigint references reception (id) NOT NULL,
    warehouse_id bigint references warehouse (id) NOT NULL
);

CREATE TABLE fermentation
(
    id           bigint PRIMARY KEY references process (id),
    created_at   timestamp                        NOT NULL,
    updated_at   timestamp,
    deleted_at   timestamp,
    warehouse_id bigint references warehouse (id) NOT NULL,
    tank_id      bigint references tank (id)      NOT NULL
);

CREATE TABLE ageing
(
    id         bigint PRIMARY KEY references process (id),
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    tank_id    bigint references tank (id) NOT NULL,
    cask_id    bigint references cask (id) NOT NULL
);

CREATE TABLE bottling
(
    id         bigint PRIMARY KEY references process (id),
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    cask_id    bigint references cask (id) NOT NULL,
    bottle_qty integer                     NOT NULL,
    wine_id    bigint references wine (id) NOT NULL
);

CREATE TABLE bottling_bottle
(
    bottling_id bigint references bottling (id),
    bottle_id   bigint references bottle (id),
    PRIMARY KEY (bottling_id, bottle_id)
);