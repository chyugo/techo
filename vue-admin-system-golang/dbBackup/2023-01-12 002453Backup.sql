/*Generated by xorm v0.7.0.0504 2023-01-12 00:24:53, from mysql to MYSQL*/

CREATE TABLE IF NOT EXISTS `heart` (`abnormal` INT(11) NOT NULL COMMENT '比率扩大10倍', `id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, `normal` INT(11) NOT NULL COMMENT '比率扩大10倍', `record_date` BIGINT(20) NOT NULL COMMENT '记录日期', `user_id` INT(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `heart` (`abnormal`, `id`, `normal`, `record_date`, `user_id`) VALUES (0, 5013, 1000, 1673398800, 1);

CREATE TABLE IF NOT EXISTS `phone` (`id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, `phone_minute` INT(11) NOT NULL, `record_date` BIGINT(15) NOT NULL COMMENT '记录日期', `user_id` INT(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `phone` (`id`, `phone_minute`, `record_date`, `user_id`) VALUES (182549, 22, 1673398800, 1);

CREATE TABLE IF NOT EXISTS `shit` (`id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, `record_date` BIGINT(20) NOT NULL COMMENT '记录日期', `shit_time` BIGINT(20) NOT NULL COMMENT '拉屎时间', `user_id` INT(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `shit` (`id`, `record_date`, `shit_time`, `user_id`) VALUES (172083, 1673312400, 1673318460, 1);

CREATE TABLE IF NOT EXISTS `sleep` (`active` INT(11) NULL COMMENT '睡眠时长分钟数', `get_time` BIGINT(20) NULL COMMENT '起床时间', `id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, `record_date` BIGINT(20) NOT NULL COMMENT '记录日期', `sleep_time` BIGINT(20) NULL COMMENT '入睡时间', `user_id` INT(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `sleep` (`active`, `get_time`, `id`, `record_date`, `sleep_time`, `user_id`) VALUES (0, 0, 309271, 1673398800, 1673427300, 1);
INSERT INTO `sleep` (`active`, `get_time`, `id`, `record_date`, `sleep_time`, `user_id`) VALUES (420, 0, 309272, 1673312400, 0, 1);
INSERT INTO `sleep` (`active`, `get_time`, `id`, `record_date`, `sleep_time`, `user_id`) VALUES (25, 1673408220, 309273, 1673398800, 1673408220, 3);

CREATE TABLE IF NOT EXISTS `suggest` (`content` VARCHAR(255) NULL, `id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `suggest` (`content`, `id`) VALUES ('12345', 1);
INSERT INTO `suggest` (`content`, `id`) VALUES ('90', 2);

CREATE TABLE IF NOT EXISTS `user` (`del_flag` INT(1) DEFAULT 0 NOT NULL, `id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, `pass_word` VARCHAR(255) NOT NULL, `user_name` VARCHAR(255) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `user` (`del_flag`, `id`, `pass_word`, `user_name`) VALUES (0, 1, '4ee71d4245f93d6d40d961a820c53966', 'chyugo');
INSERT INTO `user` (`del_flag`, `id`, `pass_word`, `user_name`) VALUES (0, 2, '905b57efc3369c05a91ae00b3b75b93c', '12');
INSERT INTO `user` (`del_flag`, `id`, `pass_word`, `user_name`) VALUES (0, 3, '4ee71d4245f93d6d40d961a820c53966', '123');

CREATE TABLE IF NOT EXISTS `weight` (`id` INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL, `morning_weight` INT(11) NULL COMMENT '早上体重扩大10倍', `night_weight` INT(11) NULL COMMENT '晚上体重扩大10倍', `noon_weight` INT(11) NULL COMMENT '中午体重扩大10倍', `record_date` BIGINT(20) NOT NULL COMMENT '记录日期', `user_id` INT(11) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `weight` (`id`, `morning_weight`, `night_weight`, `noon_weight`, `record_date`, `user_id`) VALUES (255543, 1465, 0, 0, 1673398800, 1);
INSERT INTO `weight` (`id`, `morning_weight`, `night_weight`, `noon_weight`, `record_date`, `user_id`) VALUES (255544, 1310, 0, 0, 1673398800, 3);