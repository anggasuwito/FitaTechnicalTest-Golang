/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `user_id` varchar(255) DEFAULT NULL,
  `item_sku` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `qty` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `item`;
CREATE TABLE `item` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `sku` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `qty` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `user_id` varchar(255) DEFAULT NULL,
  `total_price` double DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `transaction_history`;
CREATE TABLE `transaction_history` (
  `id` varchar(36) NOT NULL DEFAULT (uuid()),
  `transaction_id` varchar(36) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `cart_id` varchar(36) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `cart` (`id`, `user_id`, `item_sku`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('4f6ef3c4-cc3f-11ed-9138-00ffc2ef2645', 'angga123', '234234', 2, '2023-03-27 08:33:00', NULL, '2023-03-27 19:41:23');
INSERT INTO `cart` (`id`, `user_id`, `item_sku`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('6c1ba9aa-cc3b-11ed-9138-00ffc2ef2645', 'angga123', '43N23P', 1, '2023-03-27 08:05:10', NULL, '2023-03-27 19:41:23');
INSERT INTO `cart` (`id`, `user_id`, `item_sku`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('b9475547-cc91-11ed-9138-00ffc2ef2645', 'angga123', '120P90', 4, '2023-03-27 18:22:56', NULL, '2023-03-27 19:41:23');
INSERT INTO `cart` (`id`, `user_id`, `item_sku`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('ee8ce8f5-cc92-11ed-9138-00ffc2ef2645', 'angga123', 'A304SD', 4, '2023-03-27 18:31:35', NULL, '2023-03-27 19:41:23');

INSERT INTO `item` (`id`, `sku`, `name`, `price`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('7e3e036f-cc31-11ed-9138-00ffc2ef2645', '120P90', 'Google Home', 49.99, 6, '2023-03-27 06:54:06', NULL, NULL);
INSERT INTO `item` (`id`, `sku`, `name`, `price`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('7e3e451e-cc31-11ed-9138-00ffc2ef2645', '43N23P', 'MacBook Pro', 5399.99, 9, '2023-03-27 06:54:06', NULL, NULL);
INSERT INTO `item` (`id`, `sku`, `name`, `price`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('7e3e451e-cc31-11ed-9138-00ffc2ef2647', '234234', 'Raspberry Pi B', 30, 8, '2023-03-27 06:54:06', NULL, NULL);
INSERT INTO `item` (`id`, `sku`, `name`, `price`, `qty`, `created_at`, `updated_at`, `deleted_at`) VALUES
('7e3e5be4-cc31-11ed-9138-00ffc2ef2645', 'A304SD', 'Alexa Speaker', 109.5, 6, '2023-03-27 06:54:06', NULL, NULL);

INSERT INTO `transaction` (`id`, `user_id`, `total_price`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
('e4f98cb8-65ea-4246-8a85-600a2d7dc38a', 'angga123', 5974.16, 'success', '2023-03-27 19:41:23', NULL, NULL);


INSERT INTO `transaction_history` (`id`, `transaction_id`, `user_id`, `cart_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('aec61057-cc9c-11ed-9138-00ffc2ef2645', 'e4f98cb8-65ea-4246-8a85-600a2d7dc38a', 'angga123', 'ee8ce8f5-cc92-11ed-9138-00ffc2ef2645', '2023-03-27 19:41:23', NULL, NULL);
INSERT INTO `transaction_history` (`id`, `transaction_id`, `user_id`, `cart_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('aec69de9-cc9c-11ed-9138-00ffc2ef2645', 'e4f98cb8-65ea-4246-8a85-600a2d7dc38a', 'angga123', 'b9475547-cc91-11ed-9138-00ffc2ef2645', '2023-03-27 19:41:23', NULL, NULL);
INSERT INTO `transaction_history` (`id`, `transaction_id`, `user_id`, `cart_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('aec7010b-cc9c-11ed-9138-00ffc2ef2645', 'e4f98cb8-65ea-4246-8a85-600a2d7dc38a', 'angga123', '6c1ba9aa-cc3b-11ed-9138-00ffc2ef2645', '2023-03-27 19:41:23', NULL, NULL);
INSERT INTO `transaction_history` (`id`, `transaction_id`, `user_id`, `cart_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('aec7595a-cc9c-11ed-9138-00ffc2ef2645', 'e4f98cb8-65ea-4246-8a85-600a2d7dc38a', 'angga123', '4f6ef3c4-cc3f-11ed-9138-00ffc2ef2645', '2023-03-27 19:41:23', NULL, NULL);


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;