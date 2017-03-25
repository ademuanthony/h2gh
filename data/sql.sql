-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.13-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win32
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table taxmaster.auths
CREATE TABLE IF NOT EXISTS `auths` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(265) COLLATE utf8_unicode_ci NOT NULL,
  `firstName` varchar(265) COLLATE utf8_unicode_ci NOT NULL,
  `lastName` varchar(265) COLLATE utf8_unicode_ci NOT NULL,
  `hashPassword` varchar(265) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Dumping data for table taxmaster.auths: ~0 rows (approximately)
/*!40000 ALTER TABLE `auths` DISABLE KEYS */;
INSERT INTO `auths` (`id`, `email`, `password`, `firstName`, `lastName`, `hashPassword`) VALUES
	(7, 'ademuanthony@gmail.com', '', 'Anthony', 'Ademua', '$2a$10$y/STMpNcJVBZn/30EEiSs./TPkw/uYvNhu/hJkbUkrJ9VymzFd6Mm');
/*!40000 ALTER TABLE `auths` ENABLE KEYS */;

-- Dumping structure for table taxmaster.members
CREATE TABLE IF NOT EXISTS `members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `membership_id` varchar(200) NOT NULL,
  `username` varchar(200) NOT NULL,
  `parent_id` varchar(200) NOT NULL,
  `sponsor_id` varchar(200) NOT NULL,
  `level` int(11) NOT NULL,
  `left_index` int(11) NOT NULL,
  `right_index` int(11) NOT NULL,
  `registration_pin` varchar(200) NOT NULL,
  `stage` varchar(200) NOT NULL,
  `firstname` varchar(250) DEFAULT NULL,
  `lastname` varchar(250) DEFAULT NULL,
  `profile_image` varchar(256) NOT NULL,
  `phonenumber` varchar(250) DEFAULT NULL,
  `sex` varchar(250) DEFAULT NULL,
  `dob` varchar(250) DEFAULT NULL,
  `country` varchar(200) DEFAULT NULL,
  `state` varchar(200) DEFAULT NULL,
  `city` varchar(200) DEFAULT NULL,
  `address` varchar(250) DEFAULT NULL,
  `nameofkin` varchar(250) DEFAULT NULL,
  `transaction_password` varchar(250) DEFAULT NULL,
  `nextofkinaddress` varchar(250) DEFAULT NULL,
  `kinrelationship` varchar(250) DEFAULT NULL,
  `phonenumberofkin` varchar(250) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `email_address` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.members: ~0 rows (approximately)
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` (`id`, `membership_id`, `username`, `parent_id`, `sponsor_id`, `level`, `left_index`, `right_index`, `registration_pin`, `stage`, `firstname`, `lastname`, `profile_image`, `phonenumber`, `sex`, `dob`, `country`, `state`, `city`, `address`, `nameofkin`, `transaction_password`, `nextofkinaddress`, `kinrelationship`, `phonenumberofkin`, `status`, `email_address`) VALUES
	(1, 'S01C4RXY', 'tony', '', '', 1, 1, 46, '08998989', '2', 'Ademu', 'Anthony', 'baby.jpg', '08035146243', 'Female', '2016-09-25', 'Nigeria', 'Kogi', 'Imane', 'Imane', 'Ademu Anthoy E', '1234', '20 Sora Ogumakin, Iyana-Ikpaja', 'My Bother', '08035146243', 1, 'blenyo11@gmail.com');
/*!40000 ALTER TABLE `members` ENABLE KEYS */;

-- Dumping structure for table taxmaster.public_payment
CREATE TABLE IF NOT EXISTS `public_payment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tax_type_id` int(11) DEFAULT NULL,
  `status_id` int(11) DEFAULT NULL,
  `tin` varchar(18) DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `note` varchar(256) DEFAULT NULL,
  `date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK__public_payment_tax_type` (`tax_type_id`),
  KEY `FK_public_payment_status` (`status_id`),
  CONSTRAINT `FK__public_payment_tax_type` FOREIGN KEY (`tax_type_id`) REFERENCES `tax_type` (`id`),
  CONSTRAINT `FK_public_payment_status` FOREIGN KEY (`status_id`) REFERENCES `status` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.public_payment: ~6 rows (approximately)
/*!40000 ALTER TABLE `public_payment` DISABLE KEYS */;
INSERT INTO `public_payment` (`id`, `tax_type_id`, `status_id`, `tin`, `amount`, `note`, `date`) VALUES
	(1, 1, 3, '90709779', 9809, 'hoi', '2017-03-24 06:29:13'),
	(2, 2, 3, '60495849', 58000, 'Nothing', '2017-03-24 06:55:22'),
	(3, 2, 3, '8943', 56336, 'dhdhg', '2017-03-24 06:57:51'),
	(4, 1, 3, '55-2890', 6857, 'hkfyu', '2017-03-24 07:00:14'),
	(5, 1, 3, '5645', 90009, 'hkfyu', '2017-03-24 07:01:45'),
	(6, 1, 3, '5255', 2433, 'hkfyu', '2017-03-24 07:04:53');
/*!40000 ALTER TABLE `public_payment` ENABLE KEYS */;

-- Dumping structure for table taxmaster.status
CREATE TABLE IF NOT EXISTS `status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.status: ~6 rows (approximately)
/*!40000 ALTER TABLE `status` DISABLE KEYS */;
INSERT INTO `status` (`id`, `text`) VALUES
	(1, 'Active'),
	(2, 'Inactive'),
	(3, 'Pending'),
	(4, 'Approved'),
	(5, 'Rejected'),
	(6, 'Paid');
/*!40000 ALTER TABLE `status` ENABLE KEYS */;

-- Dumping structure for table taxmaster.stores
CREATE TABLE IF NOT EXISTS `stores` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `merchant_id` int(11) NOT NULL,
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `sub_domain` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `icon` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `logo` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `banner` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `url` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `remita_account_id` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `bank_account_name` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `bank_account_number` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `theme` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `header_color` varchar(16) COLLATE utf8_unicode_ci NOT NULL,
  `footer_color` varchar(16) COLLATE utf8_unicode_ci NOT NULL,
  `bottom_color` varchar(16) COLLATE utf8_unicode_ci NOT NULL,
  `phone_number` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `short_about_test` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `address` varchar(500) COLLATE utf8_unicode_ci NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Dumping data for table taxmaster.stores: ~0 rows (approximately)
/*!40000 ALTER TABLE `stores` DISABLE KEYS */;
INSERT INTO `stores` (`id`, `merchant_id`, `name`, `sub_domain`, `icon`, `logo`, `banner`, `url`, `remita_account_id`, `bank_account_name`, `bank_account_number`, `theme`, `header_color`, `footer_color`, `bottom_color`, `phone_number`, `email`, `short_about_test`, `address`, `status`) VALUES
	(1, 1, 'Food for all', 'taxmaster', '', 'logo.png', '', 'http://taxmaster.com/', '', '', '', 'taxmaster', '#008080', '#004080', '#ff0080', '+234 708 513 7865', 'laveritas@gmail.com', 'This is the best dealerr of all kinds of goods', '20 Papa Ahafa Road Iyanoba', 0);
/*!40000 ALTER TABLE `stores` ENABLE KEYS */;

-- Dumping structure for table taxmaster.tax
CREATE TABLE IF NOT EXISTS `tax` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tax_type_id` int(11) DEFAULT NULL,
  `tax_payer_id` int(11) DEFAULT NULL,
  `comment` varchar(500) DEFAULT NULL,
  `amount` float DEFAULT NULL,
  `due_date` datetime DEFAULT NULL,
  `status_id` int(11) DEFAULT '3',
  PRIMARY KEY (`id`),
  KEY `FK_tax_tax_type` (`tax_type_id`),
  KEY `FK_tax_tax_payer` (`tax_payer_id`),
  KEY `FK_tax_status` (`status_id`),
  CONSTRAINT `FK_tax_status` FOREIGN KEY (`status_id`) REFERENCES `status` (`id`),
  CONSTRAINT `FK_tax_tax_payer` FOREIGN KEY (`tax_payer_id`) REFERENCES `tax_payer` (`id`),
  CONSTRAINT `FK_tax_tax_type` FOREIGN KEY (`tax_type_id`) REFERENCES `tax_type` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.tax: ~2 rows (approximately)
/*!40000 ALTER TABLE `tax` DISABLE KEYS */;
INSERT INTO `tax` (`id`, `tax_type_id`, `tax_payer_id`, `comment`, `amount`, `due_date`, `status_id`) VALUES
	(1, 2, 1, 'dsff', 5500, NULL, 6),
	(2, 2, 1, 'Another Tax', 5600, NULL, 3);
/*!40000 ALTER TABLE `tax` ENABLE KEYS */;

-- Dumping structure for table taxmaster.tax_payer
CREATE TABLE IF NOT EXISTS `tax_payer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tin` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `contact_name` varchar(50) DEFAULT NULL,
  `contact_email` varchar(50) DEFAULT NULL,
  `contact_phone_number` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.tax_payer: ~4 rows (approximately)
/*!40000 ALTER TABLE `tax_payer` DISABLE KEYS */;
INSERT INTO `tax_payer` (`id`, `tin`, `name`, `contact_name`, `contact_email`, `contact_phone_number`) VALUES
	(1, '0987894534', 'Solid Food', '', '', ''),
	(2, '7658909872', 'Alpha Binto', '', '', ''),
	(3, '0987894534', 'Solid Food', '', '', ''),
	(4, '1234567890', 'Jude Nnadi', NULL, NULL, NULL);
/*!40000 ALTER TABLE `tax_payer` ENABLE KEYS */;

-- Dumping structure for table taxmaster.tax_payment
CREATE TABLE IF NOT EXISTS `tax_payment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tax_id` int(11) NOT NULL,
  `amount_paid` float DEFAULT NULL,
  `date` date DEFAULT NULL,
  `status_id` int(11) NOT NULL DEFAULT '3',
  PRIMARY KEY (`id`),
  KEY `FK__tax_payment_tax` (`tax_id`),
  KEY `FK_tax_payment_status` (`status_id`),
  CONSTRAINT `FK__tax_payment_tax` FOREIGN KEY (`tax_id`) REFERENCES `tax` (`id`),
  CONSTRAINT `FK_tax_payment_status` FOREIGN KEY (`status_id`) REFERENCES `status` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.tax_payment: ~3 rows (approximately)
/*!40000 ALTER TABLE `tax_payment` DISABLE KEYS */;
INSERT INTO `tax_payment` (`id`, `tax_id`, `amount_paid`, `date`, `status_id`) VALUES
	(1, 1, 5500, '2017-02-28', 3),
	(2, 2, 5600, '2017-02-28', 3),
	(3, 1, 5500, '2017-02-28', 3);
/*!40000 ALTER TABLE `tax_payment` ENABLE KEYS */;

-- Dumping structure for table taxmaster.tax_type
CREATE TABLE IF NOT EXISTS `tax_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.tax_type: ~3 rows (approximately)
/*!40000 ALTER TABLE `tax_type` DISABLE KEYS */;
INSERT INTO `tax_type` (`id`, `name`) VALUES
	(1, 'VAT'),
	(2, 'PAYE'),
	(3, 'PTF');
/*!40000 ALTER TABLE `tax_type` ENABLE KEYS */;

-- Dumping structure for table taxmaster.user_auth
CREATE TABLE IF NOT EXISTS `user_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(128) DEFAULT NULL,
  `last_name` varchar(128) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  `password` varchar(28) DEFAULT NULL,
  `hash_password` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.user_auth: ~2 rows (approximately)
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` (`id`, `first_name`, `last_name`, `email`, `password`, `hash_password`) VALUES
	(1, 'Anthony', 'Ademu', 'a@b.c', '', '$2a$10$wdHhkMdI.M8g7z3.9icH1.0vKIXSZBrI52CUMy62LTZlqGC93PXkO'),
	(2, 'Anthony', 'Ademu', 'anthonyademu@gmail.com', '', '$2a$10$lO8O7ZRXaeUdpeO4/Ks7depJdc3nUrlPymphPVrFPclj.q7ridcpW');
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
