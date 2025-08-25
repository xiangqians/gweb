-- ------------------------
-- Table structure for user
-- ------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`       INT(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `role_id`  INT(8) UNSIGNED NOT NULL COMMENT '角色主键',
    `name`     VARCHAR(50)  NOT NULL COMMENT '用户名',
    `password` VARCHAR(100) NOT NULL COMMENT '密码',
    `del`      TINYINT UNSIGNED            DEFAULT 0 COMMENT '是否已删除，0-未删除，1-已删除',
    `add_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `upd_time` DATETIME              DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB COMMENT='用户表';

-- 密码：admin
INSERT INTO `user`(`role_id`, `name`, `password`)
VALUES (1, 'admin', '$2a$10$ZsS2bA7B7AQtIBBpW7xz3OIw3FWU0CnXX7HZMi6vBNt9ZNcA2RNGG');


-- ------------------------
-- Table structure for role
-- ------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`       INT(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`     VARCHAR(50) NOT NULL COMMENT '名称',
    `code`     VARCHAR(50) NOT NULL UNIQUE COMMENT '标识码',
    `del`      TINYINT UNSIGNED           DEFAULT 0 COMMENT '是否已删除，0-未删除，1-已删除',
    `add_time` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `upd_time` DATETIME             DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB COMMENT='角色表';

INSERT INTO `role`(`name`, `code`)
VALUES ('管理员', 'ADMIN');


-- ------------------------
-- Table structure for perm
-- ------------------------
DROP TABLE IF EXISTS `perm`;
CREATE TABLE `perm`
(
    `id`       INT(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `tag`      VARCHAR(50)       DEFAULT '' COMMENT '标签',
    `name`     VARCHAR(100)      DEFAULT '' COMMENT '名称',
    `desc`     VARCHAR(100)      DEFAULT '' COMMENT '描述',
    `method`   VARCHAR(8)        DEFAULT '' COMMENT '方法',
    `path`     VARCHAR(100)      DEFAULT '' COMMENT '路径',
    `anon`     TINYINT UNSIGNED      DEFAULT 0 COMMENT '是否允许匿名访问，0-不允许，1-允许',
    `del`      TINYINT UNSIGNED       DEFAULT 0 COMMENT '是否已删除，0-未删除，1-已删除',
    `add_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `upd_time` DATETIME          DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB COMMENT='权限表';


-- -----------------------------
-- Table structure for role_perm
-- -----------------------------
DROP TABLE IF EXISTS `role_perm`;
CREATE TABLE `role_perm`
(
    `role_id`  INT(8) UNSIGNED NOT NULL COMMENT '角色主键',
    `perm_id`  INT(8) UNSIGNED NOT NULL COMMENT '权限主键',
    `add_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    PRIMARY KEY (`role_id`, `perm_id`) USING BTREE
) ENGINE=InnoDB COMMENT='角色-权限表';


-- ------------------------
-- Table structure for dict
-- ------------------------
DROP TABLE IF EXISTS `dict`;
CREATE TABLE `dict`
(
    `id`       INT(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `type`     VARCHAR(50)       DEFAULT '' COMMENT '类型',
    `name`     VARCHAR(50)       DEFAULT '' COMMENT '名称',
    `value`    VARCHAR(50)       DEFAULT '' COMMENT '值',
    `sort`     TINYINT UNSIGNED       DEFAULT 0 COMMENT '排序',
    `add_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `upd_time` DATETIME          DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB COMMENT='字典表';
