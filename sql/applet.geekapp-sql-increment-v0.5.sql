-- MySQL Script generated by MySQL Workbench
-- Sat Apr  7 23:28:19 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

-- -----------------------------------------------------
-- Change History --------------------------------------
-- -----------------------------------------------------
-- 2018-05-08 ver0.5
--    1. Add two columns in article table.
--        `internel_fid` VARCHAR(100) NOT NULL COMMENT '内部存储的fid'
--        `internel_size` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内部存储的size'
--    2. Modify columns in article_pending table.
--          ADD COLUMN `article_url` VARCHAR(1000) NOT NULL COMMENT '文章原文url地址',
--          CHANGE `article_url` `internel_fid` VARCHAR(100);
--          MODIFY `internel_fid` VARCHAR(100) NOT NULL COMMENT '内部存储的fid';
--          ADD COLUMN `internel_size` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内部存储的size';
--          ADD COLUMN `internel_url` VARCHAR(1000) NOT NULL COMMENT '内部URL地址';
--
-- 2018-05-05 ver0.4
--    1. Add three columns in article table.
--         `content_type` VARCHAR(20) NOT NULL DEFAULT 'HTML' COMMENT '文章内容格式，枚举：HTML、MD、TXT'
--         `images` VARCHAR(4096) COMMENT '文章摘要图片，jsons数组格式';
--         `preview_layout` VARCHAR(20) NOT NULL DEFAULT 'TXT' COMMENT '预览布局，枚举：TXT(纯文字)、PIC-TXT-TB(图文混排上下)、PIC-TXT-LR(图文混排左右)';
--
-- -----------------------------------------------------

USE `applet.geekapp`;
ALTER TABLE `applet.geekapp`.`article` ADD COLUMN `internel_fid` VARCHAR(100) NOT NULL COMMENT '内部存储的fid';
ALTER TABLE `applet.geekapp`.`article` ADD COLUMN `internel_size` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内部存储的size';

ALTER TABLE `applet.geekapp`.`article_pending` ADD COLUMN `article_url` VARCHAR(1000) NOT NULL COMMENT '文章原文url地址';
ALTER TABLE `applet.geekapp`.`article_pending` CHANGE `article_url` `internel_fid` VARCHAR(100);
ALTER TABLE `applet.geekapp`.`article_pending` MODIFY `internel_fid` VARCHAR(100) NOT NULL COMMENT '内部存储的fid';
ALTER TABLE `applet.geekapp`.`article_pending` ADD COLUMN `internel_size` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内部存储的size';
ALTER TABLE `applet.geekapp`.`article_pending` ADD COLUMN `internel_url` VARCHAR(1000) NOT NULL COMMENT '内部URL地址';