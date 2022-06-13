/*
 Navicat MySQL Data Transfer

 Source Server         : tcm
 Source Server Type    : MySQL
 Source Server Version : 100327
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 100327
 File Encoding         : 65001

 Date: 13/06/2022 10:18:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `cid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论表id',
  `content` varchar(255) NOT NULL COMMENT '评论内容',
  `createtime` bigint(20) NOT NULL COMMENT '评论发布时间mm-dd',
  PRIMARY KEY (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='视频评论表';

-- ----------------------------
-- Table structure for r_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS `r_comment_reply`;
CREATE TABLE `r_comment_reply` (
  `rcrid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论回复关联表id',
  `cid` bigint(20) unsigned NOT NULL COMMENT '所属评论的id',
  `fromid` bigint(20) unsigned NOT NULL COMMENT '回复id',
  `toid` bigint(20) unsigned NOT NULL COMMENT '目的评论或回复id',
  `idtype` tinyint(3) unsigned NOT NULL COMMENT '关联类型，0时toid为评论id，1时toid为回复id',
  `isdel` tinyint(3) unsigned NOT NULL COMMENT '软删除标志',
  PRIMARY KEY (`rcrid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for r_user_follow
-- ----------------------------
DROP TABLE IF EXISTS `r_user_follow`;
CREATE TABLE `r_user_follow` (
  `rufid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户关注表id',
  `fromuid` bigint(20) unsigned NOT NULL COMMENT '粉丝用户id',
  `touid` bigint(20) unsigned NOT NULL COMMENT '被关注用户id',
  `isdel` tinyint(3) unsigned NOT NULL COMMENT '软删除标志',
  PRIMARY KEY (`rufid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户关注粉丝表';

-- ----------------------------
-- Table structure for r_user_reply
-- ----------------------------
DROP TABLE IF EXISTS `r_user_reply`;
CREATE TABLE `r_user_reply` (
  `rurid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户回复关联表',
  `uid` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `rid` bigint(20) unsigned NOT NULL COMMENT '回复id',
  PRIMARY KEY (`rurid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户回复关联表';

-- ----------------------------
-- Table structure for r_user_video
-- ----------------------------
DROP TABLE IF EXISTS `r_user_video`;
CREATE TABLE `r_user_video` (
  `ruvid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关联表id',
  `uid` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `vid` bigint(20) unsigned NOT NULL COMMENT '所发视频id',
  `isdel` tinyint(3) unsigned NOT NULL COMMENT '软删除标志',
  PRIMARY KEY (`ruvid`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='用户和发布视频关联表';

-- ----------------------------
-- Table structure for r_video_comment
-- ----------------------------
DROP TABLE IF EXISTS `r_video_comment`;
CREATE TABLE `r_video_comment` (
  `rvcid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频评论关联表id',
  `uid` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `vid` bigint(20) unsigned NOT NULL COMMENT '视频id',
  `cid` bigint(20) unsigned NOT NULL COMMENT '评论id',
  `isdel` tinyint(3) unsigned NOT NULL COMMENT '软删除标志',
  PRIMARY KEY (`rvcid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='视频评论关联表';

-- ----------------------------
-- Table structure for r_video_favorite
-- ----------------------------
DROP TABLE IF EXISTS `r_video_favorite`;
CREATE TABLE `r_video_favorite` (
  `rvfid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关联表id',
  `uid` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `vid` bigint(20) unsigned NOT NULL COMMENT '点赞视频id',
  `isdel` tinyint(3) unsigned NOT NULL COMMENT '软删除标志',
  PRIMARY KEY (`rvfid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='用户和点赞视频关联表';

-- ----------------------------
-- Table structure for reply
-- ----------------------------
DROP TABLE IF EXISTS `reply`;
CREATE TABLE `reply` (
  `rid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '回复id',
  `content` varchar(255) NOT NULL COMMENT '回复内容',
  `createtime` bigint(20) NOT NULL COMMENT '回复发布时间',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='回复表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `uid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uname` varchar(255) NOT NULL COMMENT '用户名',
  `pword` varchar(255) NOT NULL COMMENT '密码',
  `salt` varchar(255) NOT NULL COMMENT '密码盐',
  `nickname` varchar(255) NOT NULL COMMENT '用户昵称',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COMMENT='用户账户及信息表';

-- ----------------------------
-- Table structure for user_count
-- ----------------------------
DROP TABLE IF EXISTS `user_count`;
CREATE TABLE `user_count` (
  `ucid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户计数表id',
  `uid` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `followcount` int(10) unsigned NOT NULL COMMENT '关注总数',
  `followercount` int(10) unsigned NOT NULL COMMENT '粉丝总数',
  PRIMARY KEY (`ucid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='用户相关计数表';

-- ----------------------------
-- Table structure for user_token
-- ----------------------------
DROP TABLE IF EXISTS `user_token`;
CREATE TABLE `user_token` (
  `utid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'token表id',
  `uid` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `token` varchar(255) NOT NULL COMMENT 'token',
  `expiredtime` bigint(20) NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`utid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='用户token表';

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `vid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `title` varchar(255) NOT NULL COMMENT '视频标题',
  `playurl` varchar(255) NOT NULL COMMENT '视频播放地址',
  `coverurl` varchar(255) NOT NULL COMMENT '视频封面地址',
  `createtime` bigint(20) NOT NULL COMMENT '视频发布时间mm-dd',
  PRIMARY KEY (`vid`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

-- ----------------------------
-- Table structure for video_count
-- ----------------------------
DROP TABLE IF EXISTS `video_count`;
CREATE TABLE `video_count` (
  `vcid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频计数表id',
  `vid` bigint(20) unsigned NOT NULL COMMENT '视频id',
  `favoritecount` int(10) unsigned NOT NULL COMMENT '点赞总数',
  `commentcount` int(10) unsigned NOT NULL COMMENT '评论总数',
  PRIMARY KEY (`vcid`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='视频相关的计数表';

SET FOREIGN_KEY_CHECKS = 1;
