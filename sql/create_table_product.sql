CREATE TABLE `product` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `date_created` DATETIME NOT NULL,
  `last_updated` DATETIME NOT NULL,
  PRIMARY KEY (`id`));