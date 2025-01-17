-- MySQL Script generated by MySQL Workbench
-- Sat Apr  7 23:28:19 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

-- -----------------------------------------------------
-- Change History --------------------------------------
-- -----------------------------------------------------
-- 2018-06-18 ver0.6
--    1. Add article_source table
--
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
CREATE TABLE `article_source` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
  `source_name` VARCHAR(1000) NOT NULL COMMENT '文章源名称',
  `source_type` VARCHAR(255) NOT NULL COMMENT '文章源类型("FEED","WEB")目前只支持"FEED"',
  `source_url` VARCHAR(1000) NOT NULL COMMENT '文章源地址',
  `gmt_create` DATETIME NOT NULL COMMENT '记录创建时间',
  `gmt_update` DATETIME NOT NULL COMMENT '记录更新时间',
  `create_user` VARCHAR(100) DEFAULT NULL COMMENT '记录创建人，格式为前缀+创建人ID，前缀为2字节，微信为we开头',
  `update_user` VARCHAR(100) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`) )
ENGINE=InnoDB AUTO_INCREMENT=277 DEFAULT CHARSET=utf8 COMMENT='文章源信息表';
