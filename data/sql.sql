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
  `status` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_members_deleted_at` (`deleted_at`),
  KEY `idx_members_referral_id` (`referral_id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.member: ~3 rows (approximately)
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` (`id`, `created_at`, `updated_at`, `deleted_at`, `referral_id`, `first_name`, `last_name`, `phone_number`, `email`, `password`, `hash_password`, `bank_id`, `account_name`, `account_number`, `status`) VALUES
	(1, NULL, NULL, NULL, 0, 'pjioij', 'joijoj', 'ojoij', 'a@b.c', '', '$2a$10$gsUVtJT7uSM1imRAn3jasO0SlwLMiabPEFNifwbCZYgkmNsxR6MdK', 9, 'ljlji', '0789', 'Active'),
	(13, NULL, NULL, NULL, 1, 'Anthony', 'Ademu', '08039283747', 'blenyo11@gmail.com', '', '$2a$10$L6GB50CqGL3VCdLMyRKJgea8LHf/ZlAcFTjVjbX1WbDHi4wSNc56y', 3, 'Ojima Ademu', '4039405034', 'Active'),
	(18, NULL, NULL, NULL, 13, 'Ojima', 'Emmanuel', '0959495049', 'ojima@gmail.com', '', '$2a$10$L6GB50CqGL3VCdLMyRKJgea8LHf/ZlAcFTjVjbX1WbDHi4wSNc56y', 3, 'Ademu Emmanuel', '3940392039', 'Active'),
	(19, NULL, NULL, NULL, 13, 'Adejo', 'Agatha', '08049394959', 'agatha@gmail.com', '', '$2a$10$L6GB50CqGL3VCdLMyRKJgea8LHf/ZlAcFTjVjbX1WbDHi4wSNc56y', 6, 'Adejo Agatha', '2030495920', 'Active'),
	(20, NULL, NULL, NULL, 13, 'Audu', 'Victoe', '08594939409', 'audu@gmail.com', '', '$2a$10$L6GB50CqGL3VCdLMyRKJgea8LHf/ZlAcFTjVjbX1WbDHi4wSNc56y', 7, 'Audu Victor', '4384958394', 'Active'),
	(21, NULL, NULL, NULL, 13, 'Lucy', 'Gabriel', '08039392039', 'lucy@gmail.com', '', '$2a$10$L6GB50CqGL3VCdLMyRKJgea8LHf/ZlAcFTjVjbX1WbDHi4wSNc56y', 12, 'Lucy Gabriel', '4950394920', 'Active'),
	(24, NULL, NULL, NULL, 13, 'Samuel', 'Gabriel', '08049394837', 'samuel@gmail.com', '', '$2a$10$N9pka51SCxomVEYFxvKV6e76hs9.3Uoz/sQfJGcCufo9.ZTSDlYFW', 1, 'Samuel Gabriel', '9302930459', 'Active');
/*!40000 ALTER TABLE `member` ENABLE KEYS */;

-- Dumping structure for table h2gh.payment
CREATE TABLE IF NOT EXISTS `payment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `to_member_id` int(11) DEFAULT NULL,
  `from_member_id` int(11) DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `queue_id` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `penalty_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_payments_member_id` (`to_member_id`),
  KEY `idx_payments_queue_id` (`queue_id`),
  KEY `idx_payments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.payment: ~2 rows (approximately)
/*!40000 ALTER TABLE `payment` DISABLE KEYS */;
INSERT INTO `payment` (`id`, `created_at`, `deleted_at`, `to_member_id`, `from_member_id`, `amount`, `description`, `queue_id`, `status`, `penalty_time`) VALUES
	(2, '2017-04-09 03:41:32', '2017-04-09 05:41:31', 1, 13, 15000, 'Membership Rebate', 1, 'Confirmed', '2017-04-10 03:41:32'),
	(3, '2017-04-09 03:41:32', '2017-04-09 05:41:31', 1, 13, 5000, 'Referral Bonus', 2, 'Confirmed', '2017-04-10 03:41:32'),
	(4, '2017-04-09 09:05:34', '2017-04-09 11:05:33', 1, 18, 15000, 'Membership Rebate', 3, 'Confirmed', '2017-04-10 09:05:34'),
	(5, '2017-04-09 09:05:34', '2017-04-09 11:05:33', 1, 18, 5000, 'Referral Bonus', 7, 'Confirmed', '2017-04-10 09:05:34'),
	(6, '2017-04-09 09:08:15', '2017-04-09 11:08:15', 13, 19, 15000, 'Membership Rebate', 6, 'Confirmed', '2017-04-10 09:08:15'),
	(7, '2017-04-09 09:08:15', '2017-04-09 11:08:15', 13, 19, 5000, 'Referral Bonus', 11, 'Confirmed', '2017-04-10 09:08:15'),
	(8, '2017-04-09 09:13:04', '2017-04-09 11:13:03', 1, 20, 15000, 'Membership Rebate', 8, 'Confirmed', '2017-04-10 09:13:04'),
	(9, '2017-04-09 09:13:04', '2017-04-09 11:13:03', 13, 20, 5000, 'Referral Bonus', 14, 'Confirmed', '2017-04-10 09:13:04'),
	(10, '2017-04-09 10:34:00', '2017-04-09 11:34:00', 18, 21, 15000, 'Membership Rebate', 10, 'Pending', '2017-04-10 10:34:00'),
	(11, '2017-04-09 10:34:00', '2017-04-09 11:34:00', 13, 21, 5000, 'Referral Bonus', 17, 'Pending', '2017-04-10 10:34:00'),
	(14, '2017-04-09 11:33:48', '2017-04-09 12:33:47', 13, 24, 15000, 'Membership Rebate', 12, 'Pending', '2017-04-10 11:33:48'),
	(15, '2017-04-09 11:33:48', '2017-04-09 12:33:47', 1, 24, 5000, 'Referral Bonus', 7, 'Pending', '2017-04-10 11:33:48');
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=latin1;

-- Dumping data for table h2gh.queue: ~6 rows (approximately)
/*!40000 ALTER TABLE `queue` DISABLE KEYS */;
INSERT INTO `queue` (`id`, `created_at`, `updated_at`, `deleted_at`, `member_id`, `amount`, `sort_order`, `description`, `status`) VALUES
	(1, NULL, NULL, NULL, 1, 15000, 1, 'Membership Rebate', 'Paid Out'),
	(2, NULL, NULL, NULL, 1, 5000, 1, 'Referral Bonus', 'Paid Out'),
	(3, NULL, '2017-04-09 10:07:05', NULL, 1, 15000, 2, 'Membership rebate', 'Paid Out'),
	(6, NULL, '2017-04-09 10:08:49', NULL, 13, 15000, 3, 'Membership Rebate', 'Paid Out'),
	(7, NULL, '2017-04-09 09:07:07', NULL, 1, 5000, 2, 'Referall Bonus', 'Paired'),
	(8, NULL, '2017-04-09 10:25:04', NULL, 1, 15000, 4, 'Membership Rebate', 'Paid Out'),
	(10, '2017-04-09 09:07:07', NULL, NULL, 18, 15000, 6, 'Membership Rebate', 'Paired'),
	(11, '2017-04-09 08:07:07', '2017-04-09 10:08:50', NULL, 13, 5000, 3, 'Referall Bonus', 'Paid Out'),
	(12, '2017-04-09 09:08:49', NULL, NULL, 13, 15000, 7, 'Membership Rebate', 'Paired'),
	(13, '2017-04-09 10:08:50', NULL, NULL, 19, 15000, 8, 'Membership Rebate', 'Pending'),
	(14, '2017-04-09 08:08:50', '2017-04-09 10:31:40', NULL, 13, 5000, 4, 'Referall Bonus', 'Paid Out'),
	(15, '2017-04-09 10:25:04', NULL, NULL, 1, 15000, 9, 'Membership Rebate', 'Pending'),
	(16, '2017-04-09 10:31:40', NULL, NULL, 20, 15000, 10, 'Membership Rebate', 'Pending'),
	(17, '2017-04-09 09:31:40', NULL, NULL, 13, 5000, 5, 'Referall Bonus', 'Paired');
/*!40000 ALTER TABLE `queue` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
