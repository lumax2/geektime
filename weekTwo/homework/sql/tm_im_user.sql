/*
 Navicat Premium Data Transfer
 File Encoding         : 65001

 Date: 31/10/2021 12:56:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tm_im_user
-- ----------------------------
DROP TABLE IF EXISTS `tm_im_user`;
CREATE TABLE `tm_im_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_uuid` varchar(200) DEFAULT NULL COMMENT '用户uuid',
  `user_name` varchar(200) DEFAULT NULL COMMENT '用户名（昵称）',
  `mobile` varchar(200) DEFAULT NULL COMMENT '手机号（base64编码）',
  `user_tag` varchar(200) DEFAULT NULL COMMENT '用户标签',
  `user_tag_code` varchar(200) DEFAULT '50011003' COMMENT '用户标签编号',
  `is_del` int(11) DEFAULT '0' COMMENT '逻辑删除标志 0-未删除 1-已删除',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建人MaxusID',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新人MaxusID',
  PRIMARY KEY (`id`),
  KEY `idx_mobile` (`mobile`),
  KEY `idx_uuid` (`user_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=616628400554291201 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
