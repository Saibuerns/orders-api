CREATE TABLE `order` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `client_id` bigint NOT NULL,
  `deliver_date` datetime NOT NULL,
  `deliver_address_id` bigint NOT NULL,
  `date_created` datetime NOT NULL,
  `last_updated` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_client_date` (`client_id`,`deliver_date`)
);