CREATE DATABASE IF NOT EXISTS `catfacts`;

USE `catfacts`;

CREATE TABLE IF NOT EXISTS `catfact` (
  id INT(11) NOT NULL AUTO_INCREMENT,
  fact VARCHAR(255) DEFAULT NULL,
  source_name VARCHAR(255) DEFAULT NULL,
  source_url VARCHAR(255) DEFAULT NULL,
  created_on DATE DEFAULT NULL,
  updated_on DATE DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB
