/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50623
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2020-02-05 11:33:50
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `content_detail`;
-- ----------------------------
DROP TABLE IF EXISTS `content_detail`;
CREATE TABLE `content_detail` (
  `id` auto_increment primary key varchar(255) NOT NULL,
  `cl_id` varchar(255) DEFAULT NULL COMMENT '内容列表id',
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户id',
  `creat_time` datetime DEFAULT NULL COMMENT '创建时间',
  `video_url` varchar(255) DEFAULT NULL COMMENT '视频路径'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of contentdetail
-- ----------------------------

-- ----------------------------
-- Table structure for `contentlist`
-- ----------------------------
DROP TABLE IF EXISTS `content_list`;
CREATE TABLE `content_list` (
  `id` varchar(255) NOT NULL,
  `lt_id` varchar(255) DEFAULT NULL COMMENT '标签id',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `imgurl` varchar(255) DEFAULT NULL COMMENT '图片路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of contentlist
-- ----------------------------

-- ----------------------------
-- Table structure for `label_table`;
-- ----------------------------
DROP TABLE IF EXISTS `label_table`;
CREATE TABLE `router_table` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of router_table
-- ----------------------------
