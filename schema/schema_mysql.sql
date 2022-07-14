CREATE TABLE process
(
    `id`          bigint                          NOT NULL AUTO_INCREMENT,
    `created_at`  datetime                        NOT NULL,
    `updated_at`  datetime                        NOT NULL,
    `deleted_at`  datetime                        NOT NULL,
    `start_date`  datetime                        NOT NULL,
    `end_date`    datetime                        NOT NULL,
    `hash`        varchar(100) CHARACTER SET utf8 NULL,
    `p_type`      char(20) CHARACTER SET utf8     NOT NULL,
    `transaction` varchar(100) CHARACTER SET utf8 NULL,
    `previous_id` bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UPROCESS` (`hash`)
);

CREATE TABLE reception
(
    `id`         bigint   NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime NOT NULL,
    `weight`     smallint NOT NULL,
    `sugar`      smallint NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `IRECEPTION` FOREIGN KEY (`id`) REFERENCES `process` (`id`)
);