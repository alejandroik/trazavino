CREATE TABLE process
(
    `id`          bigint                          NOT NULL AUTO_INCREMENT,
    `start_date`  datetime                        NOT NULL,
    `end_date`    datetime                        NOT NULL,
    `hash`        varchar(100) CHARACTER SET utf8 NULL,
    `p_type`      char(20) CHARACTER SET utf8     NOT NULL,
    `transaction` varchar(100) CHARACTER SET utf8 NOT NULL,
    `previous_id` bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UPROCESS` (`hash`)
);

CREATE TABLE reception
(
    `id`     bigint   NOT NULL,
    `weight` smallint NOT NULL,
    `sugar`  smallint NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `IRECEPTION` FOREIGN KEY (`id`) REFERENCES `process` (`id`)
);