CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `nick_name` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `avatar` varchar(200) COLLATE utf8mb4_bin NOT NULL COMMENT '头像地址',
    `status` tinyint NOT NULL DEFAULT '0',
    `pwd` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `pwd_token` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `creator` int NOT NULL DEFAULT '0',
    `updater` int NOT NULL DEFAULT '0',
    `create_at` datetime NOT NULL,
    `update_at` datetime DEFAULT NULL,
    `del_flag` tinyint NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 关于字符集排序规则的选用：https://blog.csdn.net/ory001/article/details/113984778
