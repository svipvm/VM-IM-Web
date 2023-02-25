/*
Navicat MySQL Data Transfer

Source Server         : SQL
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : vmimw

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-02-25 08:34:58
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for im_task
-- ----------------------------
DROP TABLE IF EXISTS `im_task`;
CREATE TABLE `im_task` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `user_id` int(32) NOT NULL,
  `task_name` varchar(63) DEFAULT NULL,
  `task_conf` varchar(255) DEFAULT NULL,
  `task_state` int(8) DEFAULT '0',
  `create_data` datetime DEFAULT CURRENT_TIMESTAMP,
  `modify_data` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `im_task_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `im_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for im_user
-- ----------------------------
DROP TABLE IF EXISTS `im_user`;
CREATE TABLE `im_user` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `er_name` varchar(63) DEFAULT 'IMER',
  `login_usr` varchar(63) NOT NULL,
  `salt` varchar(63) NOT NULL,
  `login_pwd` varchar(63) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_usr` (`login_usr`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
