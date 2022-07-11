CREATE TABLE process
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `start_date`  datetime                        NOT NULL,
    `end_date`    datetime                        NOT NULL,
    `hash`        varchar(100) CHARACTER SET utf8 NOT NULL,
    `p_type`      char(20) CHARACTER SET utf8     NOT NULL,
    `transaction` varchar(100) CHARACTER SET utf8 NOT NULL,
    `previous_id` bigint(20) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UPROCESS` (`hash`)
);

CREATE TABLE reception
(
    `id`            bigint(20) NOT NULL,
    `weight`        smallint(6) NOT NULL,
    `sugar`         smallint(6) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `IRECEPTION` FOREIGN KEY (`id`) REFERENCES `process` (`id`)
);