CREATE TABLE IF NOT EXISTS bottling_bottle
(
    bottling_id uuid references bottling (id),
    bottle_id   uuid references bottle (id),
    PRIMARY KEY (bottling_id, bottle_id)
);