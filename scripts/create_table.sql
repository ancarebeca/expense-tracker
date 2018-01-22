# Create DB
CREATE DATABASE IF NOT EXISTS `expenses` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `expenses`;

# Create Table
CREATE TABLE IF NOT EXISTS `statements` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
   `transaction_date` date DEFAULT NULL,
   `transaction_type` varchar(50) DEFAULT NULL,
   `sort_code` varchar(50) DEFAULT NULL,
   `account_number` varchar(50) DEFAULT NULL,
   `transaction_description` varchar(200) DEFAULT NULL,
   `debit_amount` float DEFAULT NULL,
   `credit_amount` float DEFAULT NULL,
   `balance` float DEFAULT NULL,
   `category` varchar(100) DEFAULT NULL,
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8;