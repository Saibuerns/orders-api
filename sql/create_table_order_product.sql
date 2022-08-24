CREATE TABLE `order_product` (
  `order_id` bigint NOT NULL,
  `product_id` bigint NOT NULL,
  PRIMARY KEY (`order_id`,`product_id`),
  UNIQUE KEY `uk_order_product` (`order_id`,`product_id`)
);