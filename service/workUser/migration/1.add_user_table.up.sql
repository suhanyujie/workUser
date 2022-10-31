CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `nick_name` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `status` tinyint NOT NULL DEFAULT '0',
    `pwd` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `pwd_token` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `creator` int NOT NULL DEFAULT '0',
    `updator` int NOT NULL,
    `create_at` datetime NOT NULL,
    `update_at` datetime DEFAULT NULL,
    `del_flag` tinyint NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;