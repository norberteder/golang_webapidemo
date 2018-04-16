CREATE DATABASE  IF NOT EXISTS `contacts` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `contacts`;
CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'user'@'%';

--
-- Table structure for table `contact`
--
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstName` varchar(50) DEFAULT NULL,
  `lastName` varchar(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;