# CREATE DATABASE expenses;
# USE expenses
# CREATE TABLE `statements` (
#   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
#   `transaction_date` date DEFAULT NULL,
#   `transaction_type` varchar(50) DEFAULT NULL,
#   `sort_code` varchar(50) DEFAULT NULL,
#   `account_number` varchar(50) DEFAULT NULL,
#   `transaction_description` varchar(200) DEFAULT NULL,
#   `debit_amount` float DEFAULT NULL,
#   `credit_amount` float DEFAULT NULL,
#   `balance` float DEFAULT NULL,
#   `category` varchar(100) DEFAULT NULL,
#   PRIMARY KEY (`id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8;


#test database

# CREATE DATABASE test_expenses;

USE test_expenses;

CREATE TABLE `statements` (
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