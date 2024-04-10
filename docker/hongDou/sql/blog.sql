/*
 Navicat Premium Data Transfer

 Source Server         : 8.218.92.69
 Source Server Type    : MySQL
 Source Server Version : 50743 (5.7.43)
 Source Host           : 8.218.92.69:23306
 Source Schema         : hongDou

 Target Server Type    : MySQL
 Target Server Version : 50743 (5.7.43)
 File Encoding         : 65001

 Date: 10/04/2024 14:24:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blogphoto
-- ----------------------------
DROP TABLE IF EXISTS `blogphoto`;
CREATE TABLE `blogphoto` (
  `photo` longtext,
  `date` longtext,
  `title` longtext,
  `position` longtext,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for comment_parent
-- ----------------------------
DROP TABLE IF EXISTS `comment_parent`;
CREATE TABLE `comment_parent` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `article_id` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文章ID',
  `comment` longtext COLLATE utf8mb4_unicode_ci COMMENT '评论内容',
  `comment_time` longtext COLLATE utf8mb4_unicode_ci COMMENT '评论时间',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评论人名称',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评论人邮箱地址',
  `comment_addr` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评论人IP地址',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `article_id` (`article_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for comment_sub_two
-- ----------------------------
DROP TABLE IF EXISTS `comment_sub_two`;
CREATE TABLE `comment_sub_two` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '回复ID',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父文章ID',
  `comment` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回复内容',
  `comment_time` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回复时间',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回复人名称',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回复人邮箱',
  `comment_addr` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回复人IP地址',
  `article_id` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '关联文章ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `parent_id` (`parent_id`) USING BTREE,
  KEY `aritlce_id` (`article_id`) USING BTREE,
  CONSTRAINT `parent_id` FOREIGN KEY (`parent_id`) REFERENCES `comment_parent` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for friend_table
-- ----------------------------
DROP TABLE IF EXISTS `friend_table`;
CREATE TABLE `friend_table` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` longtext COLLATE utf8mb4_unicode_ci,
  `preface` longtext COLLATE utf8mb4_unicode_ci,
  `url` longtext COLLATE utf8mb4_unicode_ci,
  `photo` longtext COLLATE utf8mb4_unicode_ci,
  `date` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for person_table
-- ----------------------------
DROP TABLE IF EXISTS `person_table`;
CREATE TABLE `person_table` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` longtext COLLATE utf8mb4_unicode_ci,
  `preface` longtext COLLATE utf8mb4_unicode_ci,
  `photo` longtext COLLATE utf8mb4_unicode_ci,
  `tag` longtext COLLATE utf8mb4_unicode_ci,
  `createTime` longtext COLLATE utf8mb4_unicode_ci,
  `updateTime` longtext COLLATE utf8mb4_unicode_ci,
  `visits` bigint(20) unsigned DEFAULT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci,
  `appear` tinyint(1) DEFAULT NULL,
  `comment` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=895 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for sensitive_words
-- ----------------------------
DROP TABLE IF EXISTS `sensitive_words`;
CREATE TABLE `sensitive_words` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(255) DEFAULT NULL,
  `word` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4039 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tagName` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=889 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` longtext COLLATE utf8mb4_unicode_ci,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `password` longtext COLLATE utf8mb4_unicode_ci,
  `uuid` longtext COLLATE utf8mb4_unicode_ci,
  `black` tinyint(1) DEFAULT NULL,
  `role` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
