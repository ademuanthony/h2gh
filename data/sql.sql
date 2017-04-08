-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.7.17-0ubuntu0.16.04.1 - (Ubuntu)
-- Server OS:                    Linux
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for h2gh
CREATE DATABASE IF NOT EXISTS `h2gh` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `h2gh`;

-- Dumping structure for table h2gh.bank
CREATE TABLE IF NOT EXISTS `bank` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.bank: ~21 rows (approximately)
/*!40000 ALTER TABLE `bank` DISABLE KEYS */;
INSERT INTO `bank` (`id`, `name`) VALUES
	(1, 'ACCESS BANK PLC'),
	(2, 'CITIBANK NIGERIA LIMITED'),
	(3, 'DIAMOND BANK PLC'),
	(4, 'ECOBANK NIGERIA PLC'),
	(5, 'ENTERPRISE BANK'),
	(6, 'FIDELITY BANK PLC'),
	(7, 'FIRST BANK OF NIGERIA PLC'),
	(8, 'FIRST CITY MONUMENT BANK PLC'),
	(9, 'GUARANTY TRUST BANK PLC'),
	(10, 'KEY STONE BANK'),
	(11, 'MAINSTREET BANK'),
	(12, 'SKYE BANK PLC'),
	(13, 'STANBIC IBTC BANK LTD.'),
	(14, 'STANDARD CHARTERED BANK NIGERIA LTD.'),
	(15, 'STERLING BANK PLC'),
	(16, 'UNION BANK OF NIGERIA PLC'),
	(17, 'UNITED BANK FOR AFRICA PLC'),
	(18, 'UNITY  BANK PLC'),
	(19, 'WEMA BANK PLC'),
	(20, 'ZENITH BANK PLC'),
	(21, 'HERITAGE BANKING COMPANY LTD.');
/*!40000 ALTER TABLE `bank` ENABLE KEYS */;

-- Dumping structure for table h2gh.member
CREATE TABLE IF NOT EXISTS `member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `referral_id` int(11) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `phone_number` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `hash_password` varchar(255) DEFAULT NULL,
  `bank_id` int(10) unsigned DEFAULT NULL,
  `account_name` varchar(255) DEFAULT NULL,
  `account_number` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_members_deleted_at` (`deleted_at`),
  KEY `idx_members_referral_id` (`referral_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.member: ~0 rows (approximately)
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
/*!40000 ALTER TABLE `member` ENABLE KEYS */;

-- Dumping structure for table h2gh.payment
CREATE TABLE IF NOT EXISTS `payment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `member_id` int(11) DEFAULT NULL,
  `from_member_id` int(11) DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `queue_id` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `penalty_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_payments_member_id` (`member_id`),
  KEY `idx_payments_queue_id` (`queue_id`),
  KEY `idx_payments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.payment: ~0 rows (approximately)
/*!40000 ALTER TABLE `payment` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment` ENABLE KEYS */;

-- Dumping structure for table h2gh.queue
CREATE TABLE IF NOT EXISTS `queue` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `sort_order` int(11) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_queues_deleted_at` (`deleted_at`),
  KEY `idx_queues_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.queue: ~0 rows (approximately)
/*!40000 ALTER TABLE `queue` DISABLE KEYS */;
/*!40000 ALTER TABLE `queue` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;