/*
Navicat MySQL Data Transfer

Source Server         : LocalFigure
Source Server Version : 50624
Source Host           : 127.0.0.1:3306
Source Database       : taxrecd

Target Server Type    : MYSQL
Target Server Version : 50624
File Encoding         : 65001

Date: 2015-12-21 01:38:02
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tax_record_ref
-- ----------------------------
DROP TABLE IF EXISTS `tax_record_ref`;
CREATE TABLE `tax_record_ref` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `org_name` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_serial_num` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_industry` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_bus_scope` varchar(1000) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_legal` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_reg_time` int(11) DEFAULT NULL,
  `org_address` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_reg_cap` double DEFAULT NULL,
  `org_tax_office` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `org_is_export` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `tax_income1` double DEFAULT NULL,
  `tax_ex_income1` double DEFAULT NULL,
  `tax_vat1` double DEFAULT NULL,
  `tax_income_tax1` double DEFAULT NULL,
  `tax_sum1` double DEFAULT NULL,
  `tax_income2` double DEFAULT NULL,
  `tax_ex_income2` double DEFAULT NULL,
  `tax_vat2` double DEFAULT NULL,
  `tax_income_tax2` double DEFAULT NULL,
  `tax_sum2` double DEFAULT NULL,
  `tax_income3` double DEFAULT NULL,
  `tax_ex_income3` double DEFAULT NULL,
  `tax_vat3` double DEFAULT NULL,
  `tax_income_tax3` double DEFAULT NULL,
  `tax_sum3` double DEFAULT NULL,
  `stat_tax_sum` double DEFAULT NULL,
  `stat_check_time` int(11) DEFAULT NULL,
  `is_important` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `stat_year` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
