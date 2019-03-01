CREATE TABLE `food_memory`.`foods`(
    `id`        bigint(20)  unsigned    NOT NULL    AUTO_INCREMENT,
    `uuid`      varchar(255)            NOT NULL    DEFAULT '',
    `image`     mediumblob              NOT NULL,
    `comment`   varchar(255)                        DEFAULT '',
    `user_name` varchar(255)            NOT NULL    DEFAULT '',
    `user_id`   INT         unsigned    NOT NULL    AUTO_INCREMENT,
    `reply`     varchar(255)                        DEFAULT '',
    PRIMARY KEY (`id`),
    PRIMARY KEY (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;