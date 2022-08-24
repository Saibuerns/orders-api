CREATE TABLE `address` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `street_name` VARCHAR(255) NOT NULL,
  `street_number` VARCHAR(255) NOT NULL,
  `comment` VARCHAR(255) NULL,
  PRIMARY KEY (`id`));