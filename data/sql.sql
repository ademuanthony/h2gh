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

-- Dumping structure for table taxmaster.status
CREATE TABLE IF NOT EXISTS `status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.status: ~5 rows (approximately)
/*!40000 ALTER TABLE `status` DISABLE KEYS */;
INSERT INTO `status` (`id`, `text`) VALUES
	(1, 'Active'),
	(2, 'Inactive'),
	(3, 'Pending'),
	(4, 'Approved'),
	(5, 'Rejected'),
	(6, 'Paid');
/*!40000 ALTER TABLE `status` ENABLE KEYS */;

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

-- Dumping data for table taxmaster.tax: ~0 rows (approximately)
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.tax_payer: ~2 rows (approximately)
/*!40000 ALTER TABLE `tax_payer` DISABLE KEYS */;
INSERT INTO `tax_payer` (`id`, `tin`, `name`, `contact_name`, `contact_email`, `contact_phone_number`) VALUES
	(1, '0987894534', 'Solid Food', '', '', ''),
	(2, '7658909872', 'Alpha Binto', '', '', ''),
	(3, '0987894534', 'Solid Food', '', '', '');
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

-- Dumping data for table taxmaster.tax_payment: ~0 rows (approximately)
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.tax_type: ~4 rows (approximately)
/*!40000 ALTER TABLE `tax_type` DISABLE KEYS */;
INSERT INTO `tax_type` (`id`, `name`) VALUES
	(1, 'VAT'),
	(2, 'PAYE'),
	(3, 'PTF'),
	(4, '');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table taxmaster.user_auth: ~0 rows (approximately)
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` (`id`, `first_name`, `last_name`, `email`, `password`, `hash_password`) VALUES
	(1, 'Ademu', 'Anthony', 'a@b.c', '', '$2a$10$5smSfdlic5ClBtblPSy7rewjFnrkzYZhHRrJ2fvvENou8blYD3idK');
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
