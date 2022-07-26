CREATE TABLE IF NOT EXISTS winery
(
    id         uuid PRIMARY KEY,
    created_at timestamp   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    name       varchar(20) NOT NULL
);